package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

var filePath = os.Args[1]

// GuessType infers the column type based on regexes
func GuessType(s string) string {
	colType := Parse(s)

	// integersOnly, err := regexp.MatchString(`^\d*$`, s)
	// dateOnly, err := regexp.MatchString(`\d{4}-\d{2}-\d{2}$`, s)
	// timeStampWithoutTZ, err := regexp.MatchString(`\d{4}-\d{2}-|\d{2}\s\d{2}\:\d{2}\:\d{2}`, s)
	// // fmt.Println("integersOnly", integersOnly, s)
	// // fmt.Println("dateOnly", dateOnly, s)
	// // fmt.Println("timestampWithoutTZ", timeStampWithoutTZ, s)
	// if err == nil {
	// 	if integersOnly {
	// 		return "INT"
	// 	} else if dateOnly {
	// 		return "DATE"
	// 	} else if timeStampWithoutTZ {
	// 		return "TIMESTAMP"
	// 	}
	// }
	return colType
}

func getTableName(filename string) string {
	name := strings.Split(path.Clean(filename), ".")
	return strings.ToLower(name[0])
}

func columnParse(c string) string {
	var splitter = make([]string, 0)
	if strings.Contains(c, "-") {
		splitter = strings.Split(c, "-")
	} else if strings.Contains(c, " ") {
		splitter = strings.Split(c, " ")
	} else {
		return c
	}

	return strings.Join(splitter, "_")

}

func buildTable(fieldTypes []map[string]string) string {
	sqlTables := make([]string, 0)
	sqlTables = append(sqlTables, "CREATE TABLE")
	sqlTables = append(sqlTables, getTableName(filePath)+"(")

	for i, f := range fieldTypes {
		for k, v := range f {
			if i < len(fieldTypes)-1 {
				sqlTables = append(sqlTables, columnParse(k)+" "+v+",")
			} else {
				sqlTables = append(sqlTables, columnParse(k)+" "+v)
			}
		}
	}

	sqlTables = append(sqlTables, ");\n")
	return strings.Join(sqlTables, " ")
}

func buildInsert(fieldValues []string) string {
	var sqlRow = "INSERT INTO " + getTableName(filePath) + " VALUES " + "("
	for i, value := range fieldValues {
		if i < len(fieldValues)-1 {
			sqlRow += "'" + value + "'" + ", "
		} else {
			sqlRow += "'" + value + "'"
		}
	}
	sqlRow += " );\n"

	return sqlRow
}

func writeToFile(row string, file *os.File) {
	file.WriteString(row)
}

// GuessFieldTypes runs GuessType on each item in slice
func GuessFieldTypes(headers []string, fields []string) []map[string]string {
	ft := make([]map[string]string, 0)

	for i := range fields {
		fieldType := GuessType(fields[i])
		fm := make(map[string]string)
		fm[headers[i]] = fieldType
		ft = append(ft, fm)
	}
	return ft
}

func readLineIntoString(f string, ch chan string) (err error) {
	fmt.Println("reading file:", f)

	file, err := os.Open(f)
	defer file.Close()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool

		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			if !isPrefix {
				break
			}

			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		line := buffer.String()

		ch <- line
	}

	if err != io.EOF {
		fmt.Printf(" > Failed: %v \n", err)
	}

	close(ch)

	return
}

func parseLineIntoValues(line string, ind int) string {
	var row string

	if ind == 1 {
		vals := strings.Split(line, "\n")
		headers := strings.Split(vals[0], ",")
		firstRow := strings.Split(vals[1], ",")
		fieldTypes := GuessFieldTypes(headers, firstRow)
		row = buildTable(fieldTypes) + buildInsert(firstRow)
	}
	if ind > 1 {
		vals := strings.Split(line, ",")
		row = buildInsert(vals)
	}
	return row
}

func main() {
	start := time.Now()
	lineChan := make(chan string, 1000)
	ind := 0
	var headers string
	var row string

	go readLineIntoString(filePath, lineChan)

	sqlFile, err := os.Create(getTableName(filePath) + ".sql")

	if err == nil {
		for line := range lineChan {
			if ind < 2 {
				headers += line + "\n"
				if ind == 1 {
					row = parseLineIntoValues(headers, ind)
				}
			} else {
				row = parseLineIntoValues(line, ind)
			}
			writeToFile(row, sqlFile)
			ind++
		}
	}

	end := time.Now()
	fmt.Printf("Wrote %v lines to file in %v\n", ind+1, end.Sub(start))
}

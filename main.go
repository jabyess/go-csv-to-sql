package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strings"
)

var filePath = os.Args[1]
var fields [][]string
var fieldTypes = make([]map[string]string, 0)
var sqlTables = make([]string, 0)

// GuessType infers the column type based on regexes
func GuessType(s string) string {
	integersOnly, err := regexp.MatchString(`^\d*$`, s)
	dateOnly, err := regexp.MatchString(`\d{4}-\d{2}-\d{2}$`, s)
	timeStampWithoutTZ, err := regexp.MatchString(`\d{4}-\d{2}-|\d{2}\s\d{2}\:\d{2}\:\d{2}`, s)
	// fmt.Println("integersOnly", integersOnly, s)
	// fmt.Println("dateOnly", dateOnly, s)
	// fmt.Println("timestampWithoutTZ", timeStampWithoutTZ, s)
	if err == nil {
		// true if it contains A-Za-z
		// false if it contains only numbers
		// if false, make integer type
		// if true, TEXT
		if integersOnly {
			return "INT"
		} else if dateOnly {
			return "DATE"
		} else if timeStampWithoutTZ {
			return "TIMESTAMP"
		}
	}
	return "TEXT"
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

func buildWholeFile(fields [][]string, fieldTypes []map[string]string) {

	sqlTableString := buildTable(fieldTypes)
	var sqlRowString string

	for i, value := range fields {
		if i > 0 {
			sqlRowString += buildInsert(value)
		}
	}

	fmt.Println(filePath)
	sqlFile, err := os.Create(getTableName(filePath) + ".sql")
	fmt.Printf("%v\n", sqlFile)
	if err == nil {
		sqlFile.WriteString(sqlTableString)
		sqlFile.WriteString(sqlRowString)
	}
}

// GuessFieldTypes runs GuessType
func GuessFieldTypes(fields []string) []map[string]string {
	ft := make([]map[string]string, 0)
	for i := range fields {
		fieldType := GuessType(fields[i])
		fm := make(map[string]string)
		fm[fields[i]] = fieldType
		ft = append(ft, fm)
	}
	return ft
}

func readLineIntoString(f string) (err error) {
	fmt.Println("reading file:", f)

	file, err := os.Open(f)
	defer file.Close()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	// var firstLine = true

	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool
		var i = 0

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
		lineWords := strings.Split(line, ",")

		fields = append(fields, lineWords)
		i++
	}

	fieldTypes = GuessFieldTypes(fields[0])

	fmt.Println(fieldTypes)

	if err != io.EOF {
		fmt.Printf(" > Failed: %v \n", err)
	}

	return
}

func main() {
	readLineIntoString(filePath)
	buildWholeFile(fields, fieldTypes)
}

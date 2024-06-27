package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func BuildTable(fieldTypes []map[string]string) string {
	sqlTables := make([]string, 0)
	sqlTables = append(sqlTables, "CREATE TABLE")
	sqlTables = append(sqlTables, GetTableName(filePath)+"(")

	for i, f := range fieldTypes {
		for k, v := range f {
			if i < len(fieldTypes)-1 {
				sqlTables = append(sqlTables, ColumnParse(k)+" "+v+",")
			} else {
				sqlTables = append(sqlTables, ColumnParse(k)+" "+v)
			}
		}
	}

	sqlTables = append(sqlTables, ");\n")
	return strings.Join(sqlTables, " ")
}

func BuildInsert(fieldValues []string) string {
	var sqlRow = "INSERT INTO " + GetTableName(filePath) + " VALUES " + "("
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

func ColumnParse(c string) string {
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

func GetTableName(filename string) string {
	name := strings.Split(path.Clean(filename), ".")
	return strings.ToLower(name[0])
}

// GuessFieldTypes runs GuessType on each item in slice
func GuessFieldTypes(headers []string, fields []string) []map[string]string {
	fieldTypes := make([]map[string]string, 0)

	for i := range fields {
		guessedType := Parse(fields[i])
		fieldMapping := make(map[string]string)
		fieldMapping[headers[i]] = guessedType
		fieldTypes = append(fieldTypes, fieldMapping)
	}
	return fieldTypes
}

func ParseLineIntoValues(line string, lineIndex int) string {
	var row string
	fmt.Println(lineIndex, line)

	if lineIndex == 1 {
		vals := strings.Split(line, "\n")
		headers := strings.Split(vals[0], ",")
		firstRow := strings.Split(vals[1], ",")
		fieldTypes := GuessFieldTypes(headers, firstRow)
		row = BuildTable(fieldTypes) + BuildInsert(firstRow)
	}
	if lineIndex > 1 {
		vals := strings.Split(line, ",")
		row = BuildInsert(vals)
	}
	return row
}

func ReadLineIntoString(f string, ch chan string) (err error) {
	file, err := os.Open(f)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		ch <- line
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	close(ch)
	return
}

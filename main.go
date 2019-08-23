package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var filePath = os.Args[1]
var fields [][]string
var fieldTypes = make([]map[string]string, 0)
var sqlTables = make([]string, 0)

func guessType(s string) string {
	m, err := regexp.MatchString(`[[:alpha:]]`, s)
	fmt.Println("checking string: ", s)
	if err == nil {
		fmt.Println(m)
		// true if it contains A-Za-z
		// false if it contains only numbers
		// if false, make integer type
		// if true, TEXT
		if m == false {
			return "INT"
		}
	}
	return "TEXT"
}

func getTableName(filename string) string {
	name := strings.Split(filename, ".")
	return strings.ToLower(name[0])
}

func buildTable(fieldTypes []map[string]string) string {
	sqlTables = append(sqlTables, "CREATE TABLE")
	sqlTables = append(sqlTables, getTableName(filePath)+"(")

	for i, f := range fieldTypes {
		fmt.Printf("%v\n", f)
		for k, v := range f {
			if i < len(fieldTypes)-1 {
				sqlTables = append(sqlTables, k+" "+v+",")
			} else {
				sqlTables = append(sqlTables, k+" "+v)
			}
		}
	}

	sqlTables = append(sqlTables, ");\n")
	fmt.Printf("%T %v\n", sqlTables, sqlTables)
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

	fmt.Println(sqlTableString)
	fmt.Println(sqlRowString)

	sqlFile, err := os.Create(getTableName(filePath) + ".sql")
	if err == nil {
		sqlFile.WriteString(sqlTableString)
		sqlFile.WriteString(sqlRowString)
	}
}

func readLineIntoString(f string) (err error) {

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

		fmt.Printf("%v\n", lineWords)
		fields = append(fields, lineWords)
		i++
	}

	if err != io.EOF {
		fmt.Printf(" > Failed: %v \n", err)
	}

	for i := range fields[0] {
		fieldType := guessType(fields[1][i])
		fm := make(map[string]string)
		fm[fields[0][i]] = fieldType
		fieldTypes = append(fieldTypes, fm)
	}
	fmt.Println(fieldTypes)

	return
}

func main() {
	readLineIntoString(filePath)
	buildWholeFile(fields, fieldTypes)

	// print(fields[0])
}

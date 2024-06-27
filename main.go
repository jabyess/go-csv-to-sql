package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var filePath = os.Args[1]
var tableName = GetTableName(filePath)

func main() {
	start := time.Now()
	lineChan := make(chan string)
	lineIndex := 0
	var headers string

	// read line by line using a channel
	go ReadLineIntoString(filePath, lineChan)

	// create the .sql file on disk using filename of csv
	sqlFile, err := os.Create(tableName + ".sql")

	if err != nil {
		log.Fatal(err)
	}

	for line := range lineChan {
		var row string

		fmt.Println("--", lineIndex, line)
		if lineIndex == 0 {
			headers += line + "\n"
		} else if lineIndex == 1 {
			headers += line + "\n"
			row = ParseLineIntoValues(headers, lineIndex)
		} else if lineIndex > 1 {
			row = ParseLineIntoValues(line, lineIndex)
		}
		sqlFile.WriteString(row)
		lineIndex++
	}

	end := time.Now()
	fmt.Printf("Wrote %v lines to file in %v\n", lineIndex, end.Sub(start))
}

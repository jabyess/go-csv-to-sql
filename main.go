package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	filePath := os.Args[1]
	start := time.Now()
	lineChan := make(chan string)
	lineIndex := 0
	var headers string

	// read line by line using a channel
	go ReadLineIntoString(filePath, lineChan)

	// create the .sql file on disk using filename of csv
	sqlFile, err := os.Create(GetTableName(filePath) + ".sql")

	if err != nil {
		log.Fatal(err)
	}

	for line := range lineChan {
		var row string

		if lineIndex == 1 {
			headers += line + "\n"
			row = ParseLineIntoValues(headers, lineIndex)
		} else {
			row = ParseLineIntoValues(line, lineIndex)
		}
		sqlFile.WriteString(row)
		lineIndex++
	}

	end := time.Now()
	fmt.Printf("Wrote %v lines to file in %v\n", lineIndex, end.Sub(start))
}

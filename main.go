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

var fields [][]string
var fieldTypes = make([]map[string]string, 0)

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
		// lineSlice = append(lineSlice, line)
		fields = append(fields, lineWords)
		// fmt.Printf("read %d chars \n", len(line))
		// fmt.Println(" > > " + line)
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
	readLineIntoString("./test.csv")
	// print(fields[0])
}

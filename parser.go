package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// all month constants, short and long
// Jan Feb Mar
// January February March etc
var longMonths = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
var shortMonths = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var timeZone3 = []string{""}
var daysInMonth = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
var monthsInYear = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

// YYYY year
// MM month
// DD day
// separators: - / .
// characters A-Z a-z
// special chars????

// September is 9 chars, longest month
type matches struct {
	wordMatches [][]string
	numMatches  [][]string
}

// Parse is the main function that sets up the parser
func Parse(s string) string {
	baseSplit := strings.Split(s, "")

	var wordMatches, numMatches [][]string

	matches := matches{wordMatches, numMatches}

	wMatches, nMatches := startParse(baseSplit, matches.wordMatches, matches.numMatches)

	finalType, err := determineType(nMatches, wMatches)

	if err == nil {
		return finalType
	}

	return "TEXT"
}

func containsString(v string, c []string) bool {
	for _, item := range c {
		if v == item {
			return true
		}
	}
	return false
}

func sliceToStr(word []string) string {
	return strings.Join(word, "")
}

func sliceToInt(nums []string) int {
	str := strings.Join(nums, "")

	num, err := strconv.Atoi(str)

	if err != nil {
		return 0
	}
	return num

}

func determineType(nums [][]string, words [][]string) (string, error) {

	var numType string

	// check date timestamp tz types first

	for i, n := range nums {
		if i > 6 && n[0] == "-" {
			numType = "TIMESTAMP WITH TIME ZONE"
		} else if i > 6 && n[0] == "+" {
			numType = "TIMESTAMP WITH TIME ZONE"
		} else if i > 10 {
			match, err := regexp.MatchString(`\w`, n[0])
			if err == nil && match == true {
				numType = "TIMESTAMP WITH TIME ZONE"
			}
		}
	}

	// check timestamp, then dates
	if len(nums) > 5 {
		return "TIMESTAMP", nil
	} else if len(nums) > 1 {
		return "DATE", nil
	} else if len(nums) > 0 {
		return "int", nil
	}

	// check words

	var isDate bool
	// if words == month or year, timestamp type is valid
	if len(words) > 0 {
		for _, word := range words {

			isDate = containsString(sliceToStr(word), TimeZoneAbbreviations)

		}
	}

	if isDate && len(numType) > 0 {
		return numType, nil
	}
	// for
	// containsString()
	// if word != month or year, return text type

	return numType, nil

}

// func determineNumType(nums [][]string) (string, error) {

// 	fmt.Println("nums:", nums)

// 	for i, n := range nums {
// 		if i > 6 && n[0] == "-" {
// 			return "TIMESTAMP WITH TIME ZONE", nil
// 		} else if i > 6 && n[0] == "+" {
// 			return "TIMESTAMP WITH TIME ZONE", nil
// 		} else if i > 10 {
// 			match, err := regexp.MatchString(`\w`, n[0])
// 			if err == nil && match == true {
// 				return "TIMESTAMP WITH TIME ZONE", nil
// 			}
// 		}
// 	}

// 	if len(nums) > 5 {
// 		return "TIMESTAMP", nil
// 	} else if len(nums) > 1 {
// 		return "DATE", nil
// 	} else if len(nums) > 0 {
// 		return "int", nil
// 	}

// 	return "", errors.New("Not a valid num Type")

// }

// func determineStringType(words [][]string) (string, error) {
// 	fmt.Println("words:", words)

// 	if len(words) > 0 {
// 		return "TEXT", nil
// 	}

// 	return "", errors.New("Not a valid text type")

// }

// StartParse begins parsing
func startParse(baseSplit []string, wordMatches [][]string, numMatches [][]string) ([][]string, [][]string) {

	wordCount := 0
	numCount := 0
	var lastMatch string

	for i, char := range baseSplit {
		if i > 30 {
			break
		}
		switch char {

		case "0":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "1":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "2":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "3":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "4":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "5":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "6":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "7":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "8":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "9":
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"

		case "A":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "a":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "B":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "b":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "C":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "c":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "D":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "d":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "E":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "e":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "F":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "f":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "G":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "g":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "H":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "h":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "I":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "i":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "J":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "j":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "K":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "k":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "L":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "l":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "M":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "m":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "N":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "n":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "O":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "o":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "P":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "p":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "Q":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "q":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "R":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "r":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "S":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "s":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "T":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "t":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "U":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "u":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "V":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "v":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "W":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "w":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "X":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "x":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "Y":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "y":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "Z":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case "z":
			wordMatches = addToWordMatch(wordMatches, wordCount, char)
			lastMatch = "word"
		case " ":
			if lastMatch == "word" {
				wordMatches = addToWordMatch(wordMatches, wordCount, char)
			} else {
				numMatches = addToNumMatch(numMatches, numCount, char)
				numCount++
				lastMatch = "num"
			}

		case ":":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "+":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "-":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			numCount++
			lastMatch = "num"
		case "/":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case ".":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			numCount++
			lastMatch = "num"
		default:
			continue
		}

	}

	for _, word := range wordMatches {
		fmt.Println("word:", word)
	}
	for _, num := range numMatches {
		fmt.Println("num:", num)
	}

	return wordMatches, numMatches
}

func addToNumMatch(numMatches [][]string, index int, s string) [][]string {
	fmt.Println(index, len(numMatches), s)

	for index > len(numMatches)-1 {
		numMatches = append(numMatches, make([]string, 0))
	}

	numMatches[index] = append(numMatches[index], s)
	fmt.Println(numMatches)
	return numMatches
}

func addToWordMatch(wordMatches [][]string, index int, s string) [][]string {
	for index > len(wordMatches)-1 {
		wordMatches = append(wordMatches, make([]string, 0))
	}
	wordMatches[index] = append(wordMatches[index], s)

	return wordMatches
}

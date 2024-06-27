package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// all month constants, short and long
// Jan Feb Mar
// January February March etc
var longMonths = []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
var shortMonths = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}

// YYYY year
// MM month
// DD day
// separators: - / .
// characters A-Z a-z
// special chars????

// September is 9 chars, longest month
type matches struct {
	wordMatches [][]string
	numMatches  [][]int
}

func addToNumMatch(numMatches [][]int, index int, s string) [][]int {
	if index > len(numMatches)-1 {
		numMatches = append(numMatches, make([]int, 0))
	}
	num, err := strconv.Atoi(s)

	if err == nil {
		numMatches[index] = append(numMatches[index], num)
	} else {
		fmt.Println("skipping character:", err)
	}
	return numMatches
}

func addToWordMatch(wordMatches [][]string, index int, s string) [][]string {
	if index > len(wordMatches)-1 {
		wordMatches = append(wordMatches, make([]string, 0))
	}
	wordMatches[index] = append(wordMatches[index], s)

	return wordMatches
}

/*
Parse is the main function that sets up the parser. it takes in a single string,
in this case, one of the values from the second line of the provided csv file
and tries to infer its type
*/
func Parse(s string) string {

	var wordMatches [][]string
	var numMatches [][]int
	matches := matches{wordMatches, numMatches}

	wMatches, nMatches := startParse(s, matches.wordMatches, matches.numMatches)

	fmt.Println("word, num matches", wMatches, nMatches)

	numType, numErr := determineNumType(nMatches)
	wordType, wordErr := determineStringType(wMatches)

	if numErr == nil {
		return numType
	} else if wordErr == nil {
		return wordType
	}

	// fmt.Println(numType, wordType)
	return wordType

}

func determineNumType(nums [][]int) (string, error) {

	// fmt.Println("nums:", nums)

	if len(nums) > 3 {
		return "timestamp", nil
	} else if len(nums) > 0 {
		return "date", nil
	}

	return "", errors.New("no valid number type")
}

func determineStringType(words [][]string) (string, error) {

	if len(words) > 0 {
		// check for boolean type
		for _, word := range words {
			letters := strings.ToLower(strings.Join(word, ""))
			if letters == "false" || letters == "true" {
				return "boolean", nil
			}
		}
		// default fallback to text type
		return "text", nil
	}

	return "", errors.New("no valid word type")

}

// StartParse begins parsing
func startParse(s string, wordMatches [][]string, numMatches [][]int) ([][]string, [][]int) {

	splitWord := strings.Split(s, "")
	// wordCount := 0
	numCount := 0
	var lastMatch string

	for i, char := range splitWord {
		// why did i put 20 as a max value originally???
		if i > 20 {
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
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "a":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "B":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "b":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "C":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "c":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "D":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "d":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "E":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "e":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "F":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "f":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "G":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "g":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "H":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "h":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "I":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "i":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "J":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "j":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "K":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "k":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "L":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "l":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "M":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "m":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "N":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "n":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "O":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "o":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "P":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "p":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "Q":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "q":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "R":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "r":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "S":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "s":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "T":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "t":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "U":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "u":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "V":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "v":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "W":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "w":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "X":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "x":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "Y":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "y":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "Z":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case "z":
			wordMatches = addToWordMatch(wordMatches, numCount, char)
			lastMatch = "word"
		case " ":
			if lastMatch == "word" {
				wordMatches = addToWordMatch(wordMatches, numCount, char)
			} else {
				numCount++
				lastMatch = "num"
				numMatches = addToNumMatch(numMatches, numCount, char)
			}

		case ":":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "-":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case "/":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		case ".":
			numCount++
			numMatches = addToNumMatch(numMatches, numCount, char)
			lastMatch = "num"
		default:
			continue
		}

	}

	// for _, word := range wordMatches {
	// 	fmt.Println("word:", word)
	// }
	// for _, num := range numMatches {
	// 	fmt.Println("num:", num)
	// }

	return wordMatches, numMatches
}

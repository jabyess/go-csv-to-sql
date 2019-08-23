package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// os.Args = []string{"cmd", "test.csv"}
	// call flag.Parse() here if TestMain uses flags
	flag.Parse()

	os.Exit(m.Run())
}

func TestGuessType(t *testing.T) {
	var columnTypes = []string{"1999-01-08", "528", "0528", "value031", "plaintext"}
	// var results = []string{"DATE", "INT", "INT", "TEXT", "TEXT"}
	t.Log(columnTypes)
	// for i, v := range columnTypes {
	// 	result := GuessType(v)
	// 	t.Log(i, result)
	// 	// if result != results[i] {
	// 	// 	// t.Errorf("test failed")
	// 	// }
	// }
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		os.Exit(1)
	}

	filePath := args[0]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dat, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	str := string(dat)
	lexer := NewLexer(str)
	parser := NewParser(lexer)
	json := parser.ParseJSON()
	if json == nil {
		log.Fatal("Error parsing JSON")
	}
	fmt.Printf("%s\n", json.String())
}

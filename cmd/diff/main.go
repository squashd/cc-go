package main

import (
	"ccdiff/diff"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		err := fmt.Errorf("usage: %s file1 file2", os.Args[0])
		log.Fatal(err)
		os.Exit(1)
	}
	filePath1 := args[0]
	filePath2 := args[1]

	file1, err := os.Open(filePath1)
	if err != nil {
		log.Fatal(err)
	}

	file2, err := os.Open(filePath2)
	if err != nil {
		log.Fatal(err)
	}

	content1, err := io.ReadAll(file1)
	if err != nil {
		log.Fatal(err)
	}
	content2, err := io.ReadAll(file2)
	if err != nil {
		log.Fatal(err)
	}

	str1 := convert(content1)
	str2 := convert(content2)

	diff.Diff(str1, str2)

}

func convert(bytes []byte) []string {
	str := string(bytes)
	lines := strings.Split(str, "\n")

	strs := make([]string, 0, len(lines))
	for _, line := range lines {
		strs = append(strs, line)
	}
	return strs
}

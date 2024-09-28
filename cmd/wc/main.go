package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	file, flag, filePath, err := getFileFlagPath(args)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	if file != os.Stdin {
		defer file.Close()
	}

	a := newFileAnalyser()
	a.registerFileFunc("-c", countFileBytes)
	a.registerFileFunc("-l", countFileLines)
	a.registerFileFunc("-w", countFileWords)
	a.registerFileFunc("-m", countFileChars)
	a.registerFileFunc("default", countFileAll)

	numString, err := a.run(flag, file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s %s\n", numString, filePath)
	os.Exit(0)

}

func getFileFlagPath(args []string) (*os.File, string, string, error) {
	var filePath, flag string
	var file *os.File
	var err error

	err = fmt.Errorf("Invalid configuration")

	flag = "default"
	if len(args) == 2 {
		filePath, flag = args[1], args[0]
		file, err = os.Open(filePath)
	}

	if len(args) == 1 {
		if fileInfo, _ := os.Stdin.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
			flag = args[0]
			file = os.Stdin
			err = nil
		} else {
			filePath = args[0]
			file, err = os.Open(filePath)
		}
	}

	if len(args) == 0 {
		if fileInfo, _ := os.Stdin.Stat(); (fileInfo.Mode() & os.ModeCharDevice) == 0 {
			file = os.Stdin
			err = nil
		} else {
			err = fmt.Errorf("have to provide a file path or pipe data")
		}
	}

	return file, flag, filePath, err
}

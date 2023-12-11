package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	var fileName = "input" // Open the file for reading
	file, err := os.Open(filepath.Join(pwd, fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	acc := 0

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop over the lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		numberOnly := ""

		runes := []rune(line) // convert to runes
		length := len(runes)

		for i := 0; i < length; i++ {
			maybeInt := string(runes[i])
			_, err := strconv.Atoi(maybeInt)
			if err == nil {
				numberOnly += maybeInt
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {
			maybeInt := string(runes[i])
			_, err := strconv.Atoi(maybeInt)
			if err == nil {
				numberOnly += maybeInt
				break
			}
		}
		doubleDigit, err := strconv.Atoi(numberOnly)
		if err == nil {
			acc += doubleDigit
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	println(acc)
}

package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
)

func groupByLength(list []string) map[int][]string {
	// create an empty map to store the result
	result := make(map[int][]string)
	// loop through the list of strings
	for _, s := range list {
		// get the length of the current string
		l := len(s)
		// check if the result map already has a key with that length
		if _, ok := result[l]; ok {
			// if yes, append the current string to the existing list
			result[l] = append(result[l], s)
		} else {
			// if no, create a new list with the current string and assign it to the key
			result[l] = []string{s}
		}
	}
	// return the result map
	return result
}

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

	numberMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	numberKeys := make([]string, 0, len(numberMap))

	// Iterate over the map using range
	for k := range numberMap {
		// Append the key to the slice
		numberKeys = append(numberKeys, k)
	}

	lengthMap := groupByLength(numberKeys)

	lengthKeys := make([]int, 0, len(lengthMap))
	for k := range lengthMap {
		// Append the key to the slice
		lengthKeys = append(lengthKeys, k)
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop over the lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		firstNumber := ""

		runes := []rune(line) // convert to runes
		length := len(runes)

		firstStop := 0

	firstNumberLoop:
		for i := 0; i < length; i++ {
			maybeInt := string(runes[i])
			_, err := strconv.Atoi(maybeInt)
			if err == nil {
				firstNumber = maybeInt
				firstStop = i
				break firstNumberLoop
			}

			for _, charCount := range lengthKeys {
				if i+1 < charCount {
					continue
				}
				lookBackString := line[i-charCount+1 : i+1]
				found := slices.Contains(lengthMap[charCount], lookBackString)
				if found {
					firstNumber = strconv.Itoa(numberMap[lookBackString])
					firstStop = i
					break firstNumberLoop
				}
			}
		}

		secondNumber := ""
		for i := firstStop; i < length; i++ {
			maybeInt := string(runes[i])
			_, err := strconv.Atoi(maybeInt)
			if err == nil {
				secondNumber = maybeInt
			}

			for _, charCount := range lengthKeys {
				if i+1 < charCount {
					continue
				}
				lookBackString := line[i-charCount+1 : i+1]
				found := slices.Contains(lengthMap[charCount], lookBackString)
				if found {
					secondNumber = strconv.Itoa(numberMap[lookBackString])
				}
			}
		}

		doubleDigit, err := strconv.Atoi(firstNumber + secondNumber)
		if err == nil {
			println(doubleDigit)
			acc += doubleDigit
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	println(acc)
}

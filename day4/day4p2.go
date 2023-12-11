package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func day4p2(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	acc := 0

	winCardIndex := map[int]int{}

	id := 1

	// Loop over the lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.Split(line, ": ")[1], " | ")
		regex := regexp.MustCompile(`\s+`)
		winningNumbers := mapset.NewSet[string]()
		winningNumbers.Append(regex.Split(parts[0], -1)...)
		winningNumbers.Remove("")
		myNumbers := mapset.NewSet[string]()
		myNumbers.Append(strings.Split(parts[1], " ")...)
		myNumbers.Remove("")
		overlap := myNumbers.Intersect(winningNumbers)
		overlapLength := float64(overlap.Cardinality())
		winCardIndex[id] = int(overlapLength)
		id++
	}

	totalCard := len(winCardIndex)
	cardCountIndex := map[int]int{}
	for i := 1; i <= totalCard; i++ {
		cardCountIndex[i] = 1
	}

	for i := 1; i <= totalCard; i++ {
		winCount := winCardIndex[i]
		cardCount := cardCountIndex[i]

		for j := 1; j <= winCount; j++ {
			jCardCount, exists := cardCountIndex[i+j]
			if exists {
				cardCountIndex[i+j] = jCardCount + cardCount
			}
		}
	}

	for _, v := range cardCountIndex {
		acc += v
	}

	fmt.Printf("%+v \n", winCardIndex)
	fmt.Printf("%+v \n", cardCountIndex)
	fmt.Printf("%+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}

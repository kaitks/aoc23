package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func day5(fileName string) int {
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
		factor := overlapLength - 1
		value := int(math.Pow(2, factor))
		acc += value
		fmt.Printf("%+v ;", winningNumbers)
		fmt.Printf("%+v ;", myNumbers)
		fmt.Printf("%+v ;", overlap)
		fmt.Printf("2 ^ %v = %v \n", factor, value)
	}

	fmt.Printf("%+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}

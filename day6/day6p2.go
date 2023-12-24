package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func day6p2(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	acc := 1

	t := readInt(scanner, "Time:")
	d := readInt(scanner, "Distance:")

	acc = countPossiblePressTime(t, d)

	fmt.Printf("%+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}

func readInt(scanner *bufio.Scanner, prefix string) int {
	scanner.Scan()
	line := strings.TrimPrefix(scanner.Text(), prefix)
	str := strings.ReplaceAll(line, " ", "")
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Invalid input:", str)
		os.Exit(1)
	}
	return num
}

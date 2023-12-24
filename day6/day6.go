package day6

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func day6(fileName string) int {
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

	times := readInts(scanner, "Time:")
	distances := readInts(scanner, "Distance:")

	for i, time := range times {
		distance := distances[i]
		count := countPossiblePressTime(time, distance)
		acc = acc * count
	}

	fmt.Printf("%+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}

func countPossiblePressTime(t int, d int) int {
	delta := math.Sqrt(float64(t*t - 4*d))
	root1 := (float64(t) - delta) / 2
	root2 := (float64(t) + delta) / 2
	mi := max(int(math.Floor(root1+1)), 0)
	ma := max(int(math.Ceil(root2-1)), mi)
	count := 0
	if ma > mi {
		count = ma - mi + 1
	}
	fmt.Printf("Time: %d, Distance: %d, Min: %d, Max: %d, Count: %d\n", t, d, mi, ma, count)
	return count
}

func readInts(scanner *bufio.Scanner, prefix string) []int {
	scanner.Scan()
	line := strings.TrimPrefix(scanner.Text(), prefix)
	fields := strings.Fields(line)
	ints := make([]int, len(fields))
	for i, str := range fields {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Invalid input:", str)
			os.Exit(1)
		}
		ints[i] = num
	}
	return ints
}

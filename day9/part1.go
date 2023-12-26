package day9

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func part1(fileName string) int {
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

	acc := 0

	for scanner.Scan() {
		line := scanner.Text()
		sequence := lop.Map(strings.Fields(line), func(str string, i int) int {
			number, _ := strconv.Atoi(str)
			return number
		})
		//fmt.Printf("Sequence: %+v\n", sequence)
		var lasts []int
		allZeroCond := false
		for !allZeroCond {
			last := sequence[len(sequence)-1]
			lasts = append(lasts, last)
			nextSequence, allZero := nextSequence(sequence)
			//fmt.Printf("Sequence: %+v\n", nextSequence)
			sequence = nextSequence
			allZeroCond = allZero
		}
		last := lo.Sum(lasts)
		//fmt.Printf("Last: %+v\n\n\n", last)
		acc += last
	}

	fmt.Printf("Total: %+v\n", acc)

	return acc
}

func nextSequence(sequence []int) ([]int, bool) {
	var nextSequence []int
	allZero := true
	for i := 0; i <= len(sequence)-2; i++ {
		delta := sequence[i+1] - sequence[i]
		nextSequence = append(nextSequence, sequence[i+1]-sequence[i])
		if delta != 0 {
			allZero = false
		}
	}
	return nextSequence, allZero
}

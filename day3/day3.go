package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func day3p2(fileName string) int {
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

	NUMBERS := mapset.NewSet[string]("0", "1", "2", "3", "4", "5", "6", "7", "8", "9")
	SYMBOL := mapset.NewSet[string]("*")

	v := 1
	numbers := []Number{}

	id := 1
	symbolLocs := mapset.NewSet[Loc]()
	gearLocs := [][]Loc{}

	acc := 0

	// Loop over the lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)
		numberString := ""
		locs := []Loc{}
		for i, item := range line {
			h := i + 1
			maybeInt := string(item)
			loc := Loc{v, h}

			if NUMBERS.Contains(maybeInt) {
				numberString += maybeInt
				locs = append(locs, loc)
			}

			if !NUMBERS.Contains(maybeInt) || h == lineLength {
				if numberString != "" {
					value, _ := strconv.Atoi(numberString)
					number := Number{strconv.Itoa(id), value, locs}
					numbers = append(numbers, number)

					// reset for next number
					id++
					numberString = ""
					locs = []Loc{}
				}
			}

			if SYMBOL.Contains(maybeInt) {
				symbolLocs.Add(loc)
			}
		}
		v++
	}

	numbersIdIndex := map[string]Number{}
	numbersLocIndex := map[Loc]string{}
	numbersIds := mapset.NewSet[string]()

	for _, number := range numbers {
		for _, loc := range number.Locs {
			numbersLocIndex[loc] = number.ID
		}

		numbersIdIndex[number.ID] = number
		numbersIds.Add(number.ID)
	}

	symbolLocsSlice := symbolLocs.ToSlice()

	for _, loc := range symbolLocsSlice {
		gearLocs = append(gearLocs, []Loc{
			{loc.V - 1, loc.H - 1},
			{loc.V, loc.H - 1},
			{loc.V + 1, loc.H - 1},
			{loc.V - 1, loc.H},
			{loc.V + 1, loc.H},
			{loc.V - 1, loc.H + 1},
			{loc.V, loc.H + 1},
			{loc.V + 1, loc.H + 1},
		})
	}

	for _, locs := range gearLocs {
		gearIds := mapset.NewSet[string]()
		for _, loc := range locs {
			number, exist := numbersLocIndex[loc]
			if exist {
				gearIds.Add(number)
			}
		}
		gearIdsSlice := gearIds.ToSlice()
		if len(gearIdsSlice) == 2 {
			firstGear := numbersIdIndex[gearIdsSlice[0]]
			secondGear := numbersIdIndex[gearIdsSlice[1]]
			acc += firstGear.Value * secondGear.Value
			fmt.Printf("%+v %d %d\n", acc, firstGear.Value, secondGear.Value)
		}
	}

	fmt.Printf("%+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}

type Number struct {
	ID    string
	Value int
	Locs  []Loc
}

type Loc struct {
	V int
	H int
}

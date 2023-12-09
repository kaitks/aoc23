package main

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	fileName := "input" // Open the file for reading
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
	SYMBOL := mapset.NewSet[string]("*", "#", "+", "$", "/", "-", "&", "=", "@", "%")

	x := 0
	numbers := []Number{}

	id := 1
	blastZone := mapset.NewSet[Loc]()
	damagedIds := mapset.NewSet[string]()
	acc := 0

	// Loop over the lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)
		numberString := ""
		locs := []Loc{}
		for y, item := range line {
			maybeInt := string(item)
			loc := Loc{x, y}

			if NUMBERS.Contains(maybeInt) {
				numberString += maybeInt
				locs = append(locs, loc)
			}

			if !NUMBERS.Contains(maybeInt) || y == lineLength-1 {
				if numberString != "" {
					value, _ := strconv.Atoi(numberString)
					number := Number{strconv.Itoa(id), value, locs}
					numbers = append(numbers, number)

					// reset for next number
					id++
					numberString = ""
					locs = []Loc{}
				}

				if SYMBOL.Contains(maybeInt) {
					blastZone.Add(loc)
				}
			}
		}
		x++
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

	for _, loc := range blastZone.ToSlice() {
		blastZone.Add(Loc{loc.X - 1, loc.Y - 1})
		blastZone.Add(Loc{loc.X, loc.Y - 1})
		blastZone.Add(Loc{loc.X + 1, loc.Y - 1})
		blastZone.Add(Loc{loc.X - 1, loc.Y})
		blastZone.Add(Loc{loc.X + 1, loc.Y})
		blastZone.Add(Loc{loc.X - 1, loc.Y + 1})
		blastZone.Add(Loc{loc.X, loc.Y + 1})
		blastZone.Add(Loc{loc.X + 1, loc.Y + 1})
	}

	for _, loc := range blastZone.ToSlice() {
		damaged := numbersLocIndex[loc]
		numbersIds.Remove(damaged)
		damagedIds.Add(damaged)
	}

	sorted := damagedIds.ToSlice()
	sort.Strings(sorted)

	fmt.Printf("%+v \n", sorted)

	for _, numberId := range sorted {
		acc += numbersIdIndex[numberId].Value
	}

	fmt.Printf("%+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Number struct {
	ID    string
	Value int
	Locs  []Loc
}

type Loc struct {
	X int
	Y int
}

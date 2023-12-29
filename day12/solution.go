package day12

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func solution(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	// Create a scanner to read the file line by line
	raw, _ := os.ReadFile(filePath)
	data := string(raw)
	lines := strings.Split(data, "\n")
	var rows []Row

	for _, lineStr := range lines {
		sections := strings.Fields(lineStr)
		Onsen := []int{}
		for _, str := range strings.Split(sections[1], ",") {
			count, _ := strconv.Atoi(str)
			Onsen = append(Onsen, count)
		}
		rows = append(rows, Row{sections[0], Onsen})
	}

	acc := 0

	for _, row := range rows {
		wayToSolve := 0
		fmt.Printf("Row: %+v\n", row.Value)
		fmt.Printf("Onsen Length: %+v\n", row.Onsen)
		process(row, &wayToSolve, 0, 0, 0)
		fmt.Printf("Way To Solve: %+v\n\n", wayToSolve)
		acc += wayToSolve
	}

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	fmt.Printf("Total: %+v\n", acc)
	return acc
}

type Row struct {
	Value string
	Onsen []int
}

func process(row Row, wayToSolve *int, valueIndex int, onsenIndex int, onsenLength int) {
outerLoop:
	for i := valueIndex; i < len(row.Value); i++ {
		str := string(row.Value[i])
		if onsenIndex < len(row.Onsen) {
			onsenTarget := row.Onsen[onsenIndex]
			switch str {
			case ".":
				if onsenLength > 0 {
					if onsenLength == onsenTarget {
						process(Row{row.Value, row.Onsen}, wayToSolve, i, onsenIndex+1, 0)
						break outerLoop
					} else {
						break outerLoop
					}
				}
			case "#":
				onsenLength++
				if onsenLength > onsenTarget {
					break outerLoop
				}
			case "?":
				process(Row{replaceStringAtIndex(row.Value, i, "."), row.Onsen}, wayToSolve, i, onsenIndex, onsenLength)
				process(Row{replaceStringAtIndex(row.Value, i, "#"), row.Onsen}, wayToSolve, i, onsenIndex, onsenLength)
				break outerLoop
			}

			if i == len(row.Value)-1 {
				if onsenLength == onsenTarget && len(row.Onsen)-1 == onsenIndex {
					fmt.Printf("Solution: %+v\n", row.Value)
					*wayToSolve++
				}
			}
		} else {
			switch str {
			case "#":
				break outerLoop
			case "?":
				process(Row{replaceStringAtIndex(row.Value, i, "."), row.Onsen}, wayToSolve, i, onsenIndex, onsenLength)
				break outerLoop
			case ".":
				if i == len(row.Value)-1 {
					fmt.Printf("Solution: %+v\n", replaceStringAtIndex(row.Value, i, "."))
					*wayToSolve++
				}
			}
		}
	}
}

func replaceStringAtIndex(str string, index int, replacement string) string {
	if index >= 0 && index < len(str) {
		return str[:index] + replacement + str[index+1:] // Combine parts
	} else {
		return str // Handle invalid index
	}
}

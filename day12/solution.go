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
		process(row, &wayToSolve)
		fmt.Printf("Row: %+v\n", row.Value)
		fmt.Printf("Onsen Length: %+v\n", row.Onsen)
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

func process(row Row, wayToSolve *int) {
	onsenLength := 0
outerLoop:
	for i := 0; i < len(row.Value); i++ {
		str := string(row.Value[i])
		if len(row.Onsen) != 0 {
			onsenTarget := row.Onsen[0]
			switch str {
			case ".":
				if onsenLength > 0 {
					if onsenLength == onsenTarget {
						process(Row{row.Value[i:], row.Onsen[1:]}, wayToSolve)
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
				process(Row{replaceStringAtIndex(row.Value, i, "."), row.Onsen}, wayToSolve)
				process(Row{replaceStringAtIndex(row.Value, i, "#"), row.Onsen}, wayToSolve)
				break outerLoop
			}

			if i == len(row.Value)-1 {
				if onsenLength == onsenTarget && len(row.Onsen) == 1 {
					*wayToSolve++
				}
			}
		} else {
			if i == len(row.Value)-1 {
				if str != "#" {
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

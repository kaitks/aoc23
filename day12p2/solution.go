package day12p2

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func solution(fileName string, repeat int) int {
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
		var Onsen []int
		for _, str := range strings.Split(sections[1], ",") {
			count, _ := strconv.Atoi(str)
			Onsen = append(Onsen, count)
		}
		var ValueRepeated []string
		for i := 0; i < repeat; i++ {
			ValueRepeated = append(ValueRepeated, sections[0])
		}
		var onsenRepeated []int
		for i := 0; i < repeat; i++ {
			onsenRepeated = append(onsenRepeated, Onsen...)
		}
		rows = append(rows, Row{strings.Join(ValueRepeated, "?"), onsenRepeated})
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

func process(row Row, wayToSolve *int, valueIndex int, onsenIndex int, accOnsenLength int) {
	onsenTargetLength := lo.Sum(row.Onsen[onsenIndex:])
	rowLength := len(row.Value)
	onsenLength := len(row.Onsen)
outerLoop:
	for i := valueIndex; i < rowLength; i++ {
		if rowLength-i+1+accOnsenLength < onsenTargetLength {
			break outerLoop
		}

		str := string(row.Value[i])
		if onsenIndex < onsenLength { // onsen haven't fully matched
			onsenTarget := row.Onsen[onsenIndex]
			switch str {
			case ".":
				if accOnsenLength > 0 {
					if accOnsenLength == onsenTarget {
						process(Row{row.Value, row.Onsen}, wayToSolve, i, onsenIndex+1, 0)
						break outerLoop
					} else {
						break outerLoop
					}
				}
			case "#":
				accOnsenLength++
				if accOnsenLength > onsenTarget {
					break outerLoop
				}
			case "?":
				process(Row{replaceStringAtIndex(row.Value, i, "."), row.Onsen}, wayToSolve, i, onsenIndex, accOnsenLength)
				process(Row{replaceStringAtIndex(row.Value, i, "#"), row.Onsen}, wayToSolve, i, onsenIndex, accOnsenLength)
				break outerLoop
			}

			if i == rowLength-1 {
				if accOnsenLength == onsenTarget && onsenLength-1 == onsenIndex {
					//fmt.Printf("Solution: %+v\n", row.Value)
					*wayToSolve++
				}
			}
		} else { // onsen matched, now the remaining str should be .
			switch str {
			case "#":
				break outerLoop
			case "?":
				process(Row{replaceStringAtIndex(row.Value, i, "."), row.Onsen}, wayToSolve, i, onsenIndex, accOnsenLength)
				break outerLoop
			case ".":
				if i == rowLength-1 {
					//fmt.Printf("Solution: %+v\n", replaceStringAtIndex(row.Value, i, "."))
					*wayToSolve++
				}
			}
		}
	}
}

func replaceStringAtIndex(str string, index int, replacement string) string {
	return str[:index] + replacement + str[index+1:] // Combine parts
}

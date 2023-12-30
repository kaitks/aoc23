package day12p2

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func solution2(fileName string, repeat int) int {
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
		fmt.Printf("Row: %+v\n", row.Value)
		fmt.Printf("Onsen Length: %+v\n", row.Onsen)
		wayToSolve := process2(row, 0, 0, 0)
		fmt.Printf("Way To Solve: %+v\n\n", wayToSolve)
		acc += wayToSolve
	}

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	fmt.Printf("Total: %+v\n", acc)
	return acc
}

func process2(row Row, valueIndex int, onsenIndex int, accOnsenLength int) int {
	minRemainValueLength := lo.Sum(row.Onsen[onsenIndex:]) + max(len(row.Onsen[onsenIndex:])-1, 0)
	rowLength := len(row.Value)
	onsenLength := len(row.Onsen)
	for i := valueIndex; i < rowLength; i++ {
		if rowLength-i+1+accOnsenLength < minRemainValueLength {
			return 0
		}

		str := string(row.Value[i])
		if onsenIndex < onsenLength { // onsen haven't fully matched
			onsenTarget := row.Onsen[onsenIndex]
			switch str {
			case "?":
				return process2(Row{replaceStringAtIndex(row.Value, i, "."), row.Onsen}, i, onsenIndex, accOnsenLength) +
					process2(Row{replaceStringAtIndex(row.Value, i, "#"), row.Onsen}, i, onsenIndex, accOnsenLength)
			case "#":
				accOnsenLength++
				if accOnsenLength > onsenTarget {
					return 0
				}
			case ".":
				if accOnsenLength > 0 {
					if accOnsenLength == onsenTarget {
						return process2(Row{row.Value, row.Onsen}, i, onsenIndex+1, 0)
					} else {
						return 0
					}
				}
			}

			if i == rowLength-1 && onsenIndex == onsenLength-1 && accOnsenLength == onsenTarget {
				//fmt.Printf("Solution: %+v\n", row.Value)
				return 1
			}
		} else { // onsen matched, now the remaining str should be .
			switch str {
			case "?":
				return process2(Row{replaceStringAtIndex(row.Value, i, "."), row.Onsen}, i, onsenIndex, accOnsenLength)
			case "#":
				return 0
			case ".":
				if i == rowLength-1 {
					//fmt.Printf("Solution: %+v\n", row.Value)
					return 1
				}
			}
		}
	}
	return 0
}

package day13

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
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
	areas := lo.Map(strings.Split(data, "\n\n"), func(area string, _ int) Area {
		rows := strings.Split(area, "\n")
		return Area{area, rows, len(rows[0]), len(rows)}
	})

	acc := 0

	for _, area := range areas {
		fmt.Printf("Area:\n%+v\n", area.toString)
		result := process(area)
		//fmt.Printf("Way To Solve: %+v\n\n", wayToSolve)
		acc += result
	}

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	fmt.Printf("Total: %+v\n", acc)
	return acc
}

type Area struct {
	toString string
	Rows     []string
	HLength  int
	VLength  int
}

func process(area Area) int {
	var possibleVerticalMirrors []int
	for i := 0; i < area.HLength-1; i++ {
		possibleVerticalMirrors = append(possibleVerticalMirrors, i)
	}
	for _, row := range area.Rows {
		var tempVM []int
		for _, i := range possibleVerticalMirrors {
			isMatch := true
			for j := 0; i-j >= 0 && i+j+1 < area.HLength; j++ {
				if row[i-j] != row[i+j+1] {
					isMatch = false
					break
				}
			}
			if isMatch {
				tempVM = append(tempVM, i)
			}
		}
		possibleVerticalMirrors = tempVM
	}
	fmt.Printf("Vertical Mirror: %+v\n", possibleVerticalMirrors)

	var possibleHorizontalMirrors []int
	for i := 0; i < area.VLength-1; i++ {
		possibleHorizontalMirrors = append(possibleHorizontalMirrors, i)
	}
	for c := 0; c < area.HLength; c++ {
		var tempHM []int
		for _, i := range possibleHorizontalMirrors {
			isMatch := true
			for j := 0; i-j >= 0 && i+j+1 < area.VLength; j++ {
				if area.Rows[i-j][c] != area.Rows[i+j+1][c] {
					isMatch = false
					break
				}
			}
			if isMatch {
				tempHM = append(tempHM, i)
			}
			possibleHorizontalMirrors = tempHM
		}
	}
	fmt.Printf("Horizontal Mirror: %+v\n\n", possibleHorizontalMirrors)

	vSum := lo.Sum(lo.Map(possibleVerticalMirrors, func(m int, i int) int {
		return m + 1
	}))
	hSum := lo.Sum(lo.Map(possibleHorizontalMirrors, func(m int, i int) int {
		return (m + 1) * 100
	}))
	return vSum + hSum
}

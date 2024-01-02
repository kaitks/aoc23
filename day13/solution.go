package day13

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
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
		rowsV := []string{}
		for i := 0; i < len(rows[0]); i++ {
			rowsV = append(rowsV, strings.Join(lo.Map(rows, func(row string, _ int) string {
				return string(row[i])
			}), ""))
		}
		return Area{area, rows, rowsV, len(rows[0]), len(rows)}
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
	toString       string
	HorizontalRows []string
	VerticalRows   []string
	HLength        int
	VLength        int
}

func process(area Area) int {
	possibleVerticalMirrors := mapset.NewSet[int]()
	for i := 0; i < area.HLength-1; i++ {
		possibleVerticalMirrors.Add(i)
	}
	for _, row := range area.HorizontalRows {
		mirros := findMirror(row)
		possibleVerticalMirrors = possibleVerticalMirrors.Intersect(mirros)
	}
	fmt.Printf("Vertical Mirror: %+v\n", possibleVerticalMirrors.ToSlice())

	possibleHorizontalMirrors := mapset.NewSet[int]()
	for i := 0; i < area.VLength-1; i++ {
		possibleHorizontalMirrors.Add(i)
	}
	for _, row := range area.VerticalRows {
		mirros := findMirror(row)
		possibleHorizontalMirrors = possibleHorizontalMirrors.Intersect(mirros)
	}
	fmt.Printf("Horizontal Mirror: %+v\n", possibleHorizontalMirrors.ToSlice())

	vSum := lo.Sum(lo.Map(possibleVerticalMirrors.ToSlice(), func(m int, i int) int {
		return m + 1
	}))
	hSum := lo.Sum(lo.Map(possibleHorizontalMirrors.ToSlice(), func(m int, i int) int {
		return (m + 1) * 100
	}))
	return vSum + hSum
}

func findMirror(row string) mapset.Set[int] {
	mirros := mapset.NewSet[int]()
	for i := 0; i < len(row)-1; i++ {
		isMatch := true
		for j := 0; i-j >= 0 && i+j+1 < len(row); j++ {
			if row[i-j] != row[i+j+1] {
				isMatch = false
				break
			}
		}
		if isMatch {
			mirros.Add(i)
		}
	}
	return mirros
}

package day11

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
)

func solution(fileName string, expandRate int) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	// Create a scanner to read the file line by line
	raw, _ := os.ReadFile(filePath)
	data := string(raw)
	rows := strings.Split(data, "\n")
	var mapp [][]string

	rowLength := 0

	for _, rowStr := range rows {
		row := strings.Split(rowStr, "")
		mapp = append(mapp, row)
		rowLength = len(row)
	}

	emptyVs := make([]bool, rowLength)

	for i := 0; i < rowLength; i++ {
		emptyVs[i] = true
	}

	galaxiesSet := mapset.NewSet[Loc]()
	emptyHSet := mapset.NewSet[int]()
	emptyVSet := mapset.NewSet[int]()

	for v := 0; v < len(mapp); v++ {
		row := mapp[v]
		isEmpty := true
		for h, str := range row {
			if str == "#" {
				galaxiesSet.Add(Loc{h, v})
				isEmpty = false
				emptyVs[h] = false
			}
		}
		if isEmpty {
			emptyHSet.Add(v)
		}
	}

	for i, isEmpty := range emptyVs {
		if isEmpty {
			emptyVSet.Add(i)
		}
	}

	galaxiesSet.ToSlice()

	galaxyPairs := [][]Loc{}
	galaxies := galaxiesSet.ToSlice()
	for i := 0; i < galaxiesSet.Cardinality(); i++ {
		for j := i + 1; j < galaxiesSet.Cardinality(); j++ {
			galaxyPair := []Loc{galaxies[i], galaxies[j]}
			galaxyPairs = append(galaxyPairs, galaxyPair)
		}
	}

	emptyRows := emptyHSet.ToSlice()
	emptyColumns := emptyVSet.ToSlice()
	slices.Sort(emptyRows)
	slices.Sort(emptyColumns)

	distances := []int{}
	multiplier := expandRate - 1

	for _, galaxyPair := range galaxyPairs {
		first := galaxyPair[0]
		second := galaxyPair[1]
		minH := min(first.H, second.H)
		maxH := max(first.H, second.H)
		minV := min(first.V, second.V)
		maxV := max(first.V, second.V)
		crossH := countElementsInRange(emptyColumns, minH, maxH)
		crossV := countElementsInRange(emptyRows, minV, maxV)
		distance := maxH - minH + maxV - minV + crossH*multiplier + crossV*multiplier
		distances = append(distances, distance)
	}

	sum := lo.Sum(distances)

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	//fmt.Printf("Empty Column: %+v\n", emptyColumns)
	//fmt.Printf("Galaxy Pairs: %+v\n", galaxyPairs)
	fmt.Printf("Sum: %+v\n", sum)
	return sum
}

type Loc struct {
	H int
	V int
}

func countElementsInRange(numbers []int, X int, Y int) int {
	count := 0
	leftIndex := sort.Search(len(numbers), func(i int) bool { return numbers[i] > X })
	rightIndex := sort.Search(len(numbers), func(i int) bool { return numbers[i] >= Y })

	count = rightIndex - leftIndex
	return count
}

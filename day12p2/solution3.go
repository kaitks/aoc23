package day12p2

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

func solution3(fileName string, repeat int) int {
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
		dfs := memoizedDfs()
		wayToSolve := dfs(row.Value, row.Onsen)
		fmt.Printf("Way To Solve: %+v\n\n", wayToSolve)
		acc += wayToSolve
	}

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	fmt.Printf("Total: %+v\n", acc)
	return acc
}

func memoizedDfs() func(string, []int) int {
	cache := map[string]map[string]int{}

	var dfs func(string, []int) int
	dfs = func(sequence string, groups []int) (v int) {
		groupsHash := fmt.Sprintf("%v", groups)
		defer func() {
			if _, ok := cache[sequence][groupsHash]; !ok {
				cache[sequence] = make(map[string]int)
				cache[sequence][groupsHash] = v
			}
		}()
		if v, ok := cache[sequence][groupsHash]; ok {
			return v
		}

		if len(groups) == 0 {
			if strings.IndexByte(sequence, '#') == -1 {
				return 1
			} else {
				return 0
			} // Check for '#' not in sequence
		}

		seqLen := len(sequence)
		groupLen := groups[0]

		if seqLen-lo.Sum(groups)-len(groups)+1 < 0 {
			return 0
		}

		hasHoles := strings.ContainsAny(sequence[:groupLen], ".") // Check for holes in the group

		if seqLen == groupLen {
			if hasHoles {
				return 0
			} else {
				return 1
			}
		}

		canUse := !hasHoles && sequence[groupLen] != '#'

		if sequence[0] == '#' {
			if canUse {
				return dfs(strings.TrimLeft(sequence[groupLen+1:], "."), groups[1:])
			} else {
				return 0
			}
		}

		skip := dfs(strings.TrimLeft(sequence[1:], "."), groups)
		if !canUse {
			return skip
		}

		return skip + dfs(strings.TrimLeft(sequence[groupLen+1:], "."), groups[1:])
	}
	return dfs
}

type IntSlice []int

func (s IntSlice) equal(a interface{}) bool {
	return reflect.DeepEqual(a, s)
}

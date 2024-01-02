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
		var rowsV []string
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
	sum := 0
outerLoop:
	for v := 0; v < area.VLength; v++ {
		for h := 0; h < area.HLength; h++ {
			hRows := area.HorizontalRows
			vRows := area.VerticalRows
			if hRows[v][h] == '.' {
				hRows[v] = replaceStringAtIndex(hRows[v], h, "#")
				vRows[h] = replaceStringAtIndex(vRows[h], v, "#")
				sum = findSum(hRows, vRows)
				if sum > 0 {
					break outerLoop
				}
			}
		}
	}
	return sum
}

func findSum(hRows []string, vRows []string) int {
	findMirror := memoize(findMirror)
	possibleVerticalMirrors := mapset.NewSet[int]()
	for i := 0; i < len(hRows[0])-1; i++ {
		possibleVerticalMirrors.Add(i)
	}
	for _, row := range hRows {
		mirros := findMirror(row)
		possibleVerticalMirrors = possibleVerticalMirrors.Intersect(mirros)
		if possibleVerticalMirrors.Cardinality() == 0 {
			break
		}
	}
	if possibleVerticalMirrors.Cardinality() > 0 {
		fmt.Printf("Vertical Mirror: %+v\n", possibleVerticalMirrors.ToSlice())
	}

	possibleHorizontalMirrors := mapset.NewSet[int]()
	for i := 0; i < len(vRows[0])-1; i++ {
		possibleHorizontalMirrors.Add(i)
	}
	for _, row := range vRows {
		mirros := findMirror(row)
		possibleHorizontalMirrors = possibleHorizontalMirrors.Intersect(mirros)
		if possibleHorizontalMirrors.Cardinality() == 0 {
			break
		}
	}
	if possibleHorizontalMirrors.Cardinality() > 0 {
		fmt.Printf("Horizontal Mirror: %+v\n", possibleHorizontalMirrors.ToSlice())
	}

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

func replaceStringAtIndex(str string, index int, replacement string) string {
	if index >= 0 && index < len(str) {
		return str[:index] + replacement + str[index+1:] // Combine parts
	} else {
		return str // Handle invalid index
	}
}

func memoize(f func(string) mapset.Set[int]) func(string) mapset.Set[int] {
	cache := make(map[string]mapset.Set[int])
	return func(x string) mapset.Set[int] {
		if val, ok := cache[x]; ok {
			return val
		}
		val := f(x)
		cache[x] = val
		return val
	}
}

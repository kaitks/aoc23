package day13

import (
	"aoc23/utils"
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
		//fmt.Printf("Area:\n%+v\n", area.toString)
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
	hRows := area.HorizontalRows
	vRows := area.VerticalRows
	originVM, originHM := findSum(hRows, vRows)
	for v := 0; v < area.VLength; v++ {
		for h := 0; h < area.HLength; h++ {
			hRowsNew := make([]string, len(hRows))
			vRowsNew := make([]string, len(vRows))
			copy(hRowsNew, hRows)
			copy(vRowsNew, vRows)
			if hRows[v][h] == '.' {
				hRowsNew[v] = utils.ReplaceStringAtIndex(hRowsNew[v], h, "#")
				vRowsNew[h] = utils.ReplaceStringAtIndex(vRowsNew[h], v, "#")
			} else {
				hRowsNew[v] = utils.ReplaceStringAtIndex(hRowsNew[v], h, ".")
				vRowsNew[h] = utils.ReplaceStringAtIndex(vRowsNew[h], v, ".")
			}
			vM, hM := findSum(hRowsNew, vRowsNew)
			finalVM := vM.Difference(originVM)
			finalHM := hM.Difference(originHM)
			if finalHM.Cardinality() > 0 || finalVM.Cardinality() > 0 {
				vSum := lo.Sum(lo.Map(finalVM.ToSlice(), func(m int, i int) int {
					return m + 1
				}))
				hSum := lo.Sum(lo.Map(finalHM.ToSlice(), func(m int, i int) int {
					return (m + 1) * 100
				}))
				return vSum + hSum
			}
		}
	}
	return sum
}

func findSum(hRows []string, vRows []string) (mapset.Set[int], mapset.Set[int]) {
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
	//if possibleVerticalMirrors.Cardinality() > 0 {
	//	fmt.Printf("Horizontal Row: \n%+v\n", strings.Join(hRows, "\n"))
	//	fmt.Printf("Vertical Mirror: %+v\n", possibleVerticalMirrors.ToSlice())
	//}

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
	//if possibleHorizontalMirrors.Cardinality() > 0 {
	//	fmt.Printf("Vertical Row: \n%+v\n", strings.Join(vRows, "\n"))
	//	fmt.Printf("Horizontal Mirror: %+v\n", possibleHorizontalMirrors.ToSlice())
	//}

	return possibleVerticalMirrors, possibleHorizontalMirrors
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

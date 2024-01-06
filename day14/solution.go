package day14

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"os"
	"path/filepath"
	"slices"
	"sort"
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
	RRock := mapset.NewSet[Loc]()
	CRock := mapset.NewSet[Loc]()
	rows := strings.Split(data, "\n")
	for v, row := range rows {
		for h, tile := range strings.Split(row, "") {
			loc := Loc{h, v}
			switch tile {
			case "O":
				RRock.Add(loc)
			case "#":
				CRock.Add(loc)
			}
		}
	}

	mapp := Map{HLength: len(rows[0]), VLength: len(rows)}
	mapp.updateCRock(CRock)
	mapp.updateRRock(RRock)

	seens := []mapset.Set[Loc]{RRock}
	hasSeen := false
	loop := 1000000000
	for i := 0; i < loop; i++ {
		tilt(&mapp, "up")
		tilt(&mapp, "left")
		tilt(&mapp, "down")
		tilt(&mapp, "right")
		if !hasSeen {
			j := slices.IndexFunc(seens, func(seen mapset.Set[Loc]) bool {
				return seen.Equal(mapp.RRock)
			})
			if j != -1 {
				hasSeen = true
				fmt.Printf("Loop: %d %d\n", j, i)
				needLoop := (1000000000 - i - 1) % (i - j + 1)
				loop = i + needLoop + 1
			} else {
				seens = append(seens, mapp.RRock)
			}
		}
		//printMap(&mapp)
	}
	acc := 0

	for _, rock := range mapp.RRock.ToSlice() {
		acc += mapp.VLength - rock.V
	}

	fmt.Printf("Total: %+v\n", acc)
	return acc
}

func tilt(mapp *Map, direction string) {
	rrNew := mapset.NewSet[Loc]()
	for _, rr := range mapp.RRock.ToSlice() {
		newLoc := getNewLoc(mapp, rr, direction)
		rrNew.Add(newLoc)
	}
	mapp.updateRRock(rrNew)
}

func getNewLoc(mapp *Map, loc Loc, direction string) Loc {
	if direction == "up" {
		blocker := findNearestSmaller(mapp.CRockHIndex[loc.H], loc.V, -1)
		rrInDirection := countElementsInRange(mapp.RRockHIndex[loc.H], blocker, loc.V)
		newPos := blocker + rrInDirection + 1
		return Loc{loc.H, newPos}
	} else if direction == "down" {
		blocker := findNearestBigger(mapp.CRockHIndex[loc.H], loc.V, mapp.VLength)
		rrInDirection := countElementsInRange(mapp.RRockHIndex[loc.H], loc.V, blocker)
		newPos := blocker - rrInDirection - 1
		return Loc{loc.H, newPos}
	} else if direction == "right" {
		blocker := findNearestBigger(mapp.CRockVIndex[loc.V], loc.H, mapp.HLength)
		rrInDirection := countElementsInRange(mapp.RRockVIndex[loc.V], loc.H, blocker)
		newPos := blocker - rrInDirection - 1
		return Loc{newPos, loc.V}
	} else if direction == "left" {
		blocker := findNearestSmaller(mapp.CRockVIndex[loc.V], loc.H, -1)
		rrInDirection := countElementsInRange(mapp.RRockVIndex[loc.V], blocker, loc.H)
		newPos := blocker + rrInDirection + 1
		return Loc{newPos, loc.V}
	}
	return loc
}

type Map struct {
	HLength     int
	VLength     int
	RRock       mapset.Set[Loc]
	RRockHIndex map[int][]int
	RRockVIndex map[int][]int
	CRock       mapset.Set[Loc]
	CRockHIndex map[int][]int
	CRockVIndex map[int][]int
}

func (mapp *Map) updateCRock(CRock mapset.Set[Loc]) {
	CRockHIndex := map[int][]int{}
	CRockVIndex := map[int][]int{}
	for _, rock := range CRock.ToSlice() {
		if v, ok := CRockHIndex[rock.H]; ok {
			CRockHIndex[rock.H] = append(v, rock.V)
		} else {
			CRockHIndex[rock.H] = []int{rock.V}
		}
		if v, ok := CRockVIndex[rock.V]; ok {
			CRockVIndex[rock.V] = append(v, rock.H)
		} else {
			CRockVIndex[rock.V] = []int{rock.H}
		}
	}
	for _, v := range CRockHIndex {
		slices.Sort(v)
	}
	for _, v := range CRockVIndex {
		slices.Sort(v)
	}
	mapp.CRockHIndex = CRockHIndex
	mapp.CRockVIndex = CRockVIndex
	mapp.CRock = CRock
}

func (mapp *Map) updateRRock(RRock mapset.Set[Loc]) {
	RRockHIndex := map[int][]int{}
	RRockVIndex := map[int][]int{}

	for _, rock := range RRock.ToSlice() {
		RRockHIndex[rock.H] = append(RRockHIndex[rock.H], rock.V)
		RRockVIndex[rock.V] = append(RRockVIndex[rock.V], rock.H)
	}
	for _, v := range RRockHIndex {
		slices.Sort(v)
	}
	for _, v := range RRockVIndex {
		slices.Sort(v)
	}
	mapp.RRockHIndex = RRockHIndex
	mapp.RRockVIndex = RRockVIndex
	mapp.RRock = RRock
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

func findNearestSmaller(numbers []int, x int, notfound int) int {
	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] > x
	})
	if index == 0 {
		return notfound
	}
	return numbers[index-1]
}

func findNearestBigger(numbers []int, x int, notfound int) int {
	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] > x
	})
	if index == len(numbers) {
		return notfound
	}
	return numbers[index]
}

func printMap(mapp *Map) {
	for v := 0; v < mapp.VLength; v++ {
		for h := 0; h < mapp.HLength; h++ {
			loc := Loc{h, v}
			if mapp.RRock.Contains(loc) {
				fmt.Print("O")
			} else if mapp.CRock.Contains(loc) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

package day14

import (
	"fmt"
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
	var RRock []Loc
	var CRock []Loc
	rows := strings.Split(data, "\n")
	for v, row := range rows {
		for h, tile := range strings.Split(row, "") {
			loc := Loc{h, v}
			switch tile {
			case "O":
				RRock = append(RRock, loc)
			case "#":
				CRock = append(CRock, loc)
			}
		}
	}

	mapp := Map{HLength: len(rows[0]), VLength: len(rows)}
	mapp.updateCRock(CRock)
	mapp.updateRRock(RRock)

	for i := 0; i < 1000; i++ {
		tilt(&mapp, "up")
		//tilt(&mapp, "left")
		//tilt(&mapp, "down")
		//tilt(&mapp, "right")
	}

	acc := 0

	for _, rock := range mapp.RRock {
		acc += mapp.VLength - rock.V
	}

	//fmt.Printf("Empty Row: %+v\n", emptyRows)
	fmt.Printf("Total: %+v\n", acc)
	return acc
}

func tilt(mapp *Map, direction string) *Map {
	var rrNew []Loc
	for _, rr := range mapp.RRock {
		newLoc := getNewLoc(mapp, rr, direction)
		rrNew = append(rrNew, newLoc)
	}
	mapp.updateRRock(rrNew)
	return mapp
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
		rrInDirection := countElementsInRange(mapp.RRockVIndex[loc.V], blocker, loc.V)
		newPos := blocker + rrInDirection + 1
		return Loc{newPos, loc.V}
	}
	return loc
}

type Map struct {
	RRock       []Loc
	CRock       []Loc
	HLength     int
	VLength     int
	CRockHIndex map[int][]int
	CRockVIndex map[int][]int
	RRockHIndex map[int][]int
	RRockVIndex map[int][]int
}

func (mapp *Map) updateCRock(CRock []Loc) {
	CRockHIndex := map[int][]int{}
	CRockVIndex := map[int][]int{}
	for _, rock := range CRock {
		CRockHIndex[rock.H] = append(CRockHIndex[rock.H], rock.V)
		CRockVIndex[rock.V] = append(CRockHIndex[rock.V], rock.H)
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

func (mapp *Map) updateRRock(RRock []Loc) {
	RRockHIndex := map[int][]int{}
	RRockVIndex := map[int][]int{}

	for _, rock := range RRock {
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

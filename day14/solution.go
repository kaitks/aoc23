package day14

import (
	"fmt"
	"github.com/samber/lo"
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

	for i := 0; i < 1; i++ {
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
		crInSameDirection, _ := mapp.CRockHIndex[loc.H]
		blockerIndex := sort.Search(len(crInSameDirection), func(i int) bool {
			rock := crInSameDirection[i]
			return rock > loc.V
		})
		if blockerIndex == len(crInSameDirection) || blockerIndex == 0 {
			blockerIndex = -1
		} else {
			blockerIndex = blockerIndex - 1
		}
		blocker := -1
		if blockerIndex > -1 {
			blocker = crInSameDirection[blockerIndex]
		}
		rrInSameDirection, _ := mapp.RRockHIndex[loc.H]
		rrInDirection := countElementsInRange(rrInSameDirection, blocker, loc.V)
		return Loc{loc.H, blocker + rrInDirection + 1}
	} else if direction == "down" {
		crInDirection := lo.Filter(mapp.CRock, func(rock Loc, _ int) bool {
			return rock.H == loc.H && rock.V > loc.V
		})
		crInDirection = append(crInDirection, Loc{loc.H, mapp.VLength})
		blocker := lo.MinBy(crInDirection, func(a Loc, b Loc) bool {
			return a.V < b.V
		})
		rrInDirection := lo.Filter(mapp.RRock, func(rock Loc, _ int) bool {
			return rock.H == loc.H && rock.V > loc.V && rock.V < blocker.V
		})
		return Loc{loc.H, blocker.V - len(rrInDirection) - 1}
	} else if direction == "right" {
		crInDirection := lo.Filter(mapp.CRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && rock.H > loc.H
		})
		crInDirection = append(crInDirection, Loc{mapp.HLength, loc.V})
		blocker := lo.MinBy(crInDirection, func(a Loc, b Loc) bool {
			return a.H < b.H
		})
		rrInDirection := lo.Filter(mapp.RRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && loc.H < rock.H && rock.H < blocker.H
		})
		return Loc{blocker.H - len(rrInDirection) - 1, loc.V}
	} else if direction == "left" {
		crInDirection := lo.Filter(mapp.CRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && rock.H < loc.H
		})
		crInDirection = append(crInDirection, Loc{-1, loc.V})
		blocker := lo.MaxBy(crInDirection, func(a Loc, b Loc) bool {
			return a.H > b.H
		})
		rrInDirection := lo.Filter(mapp.RRock, func(rock Loc, _ int) bool {
			return rock.V == loc.V && loc.H > rock.H && rock.H > blocker.H
		})
		return Loc{blocker.H + len(rrInDirection) + 1, loc.V}
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
		//if _, ok := CRockHIndex[rock.H]; !ok {
		//	CRockHIndex[rock.H] = []int{}
		//}
		//if _, ok := CRockHIndex[rock.V]; !ok {
		//	CRockHIndex[rock.V] = []int{}
		//}
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
		//if _, ok := RRockHIndex[rock.H]; !ok {
		//	RRockHIndex[rock.H] = []int{}
		//}
		//if _, ok := RRockVIndex[rock.V]; !ok {
		//	RRockVIndex[rock.V] = []int{}
		//}
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

package day17

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strconv"
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
	var mappData [][]int
	for _, rowStr := range strings.Split(data, "\n") {
		row := lo.Map(strings.Split(rowStr, ""), func(str string, _ int) int {
			value, _ := strconv.Atoi(str)
			return value
		})
		mappData = append(mappData, row)
	}
	hLength := len(mappData[0])
	vLength := len(mappData)
	mapp := Map{mappData, hLength, vLength, Loc{hLength - 1, vLength - 1}}
	bestCost := 99999999
	starter := Point{Loc{-1, 0}, Loc{0, 0}}
	path := mapset.NewSet[Loc]()
	starter.findWay(&mapp, path, []Loc{}, 0, 0, 0, &bestCost)

	fmt.Printf("Total: %+v\n", bestCost)
	return bestCost
}

type Point struct {
	current Loc
	next    Loc
}

type Loc struct {
	h int
	v int
}

func (point *Point) findWay(mapp *Map, path mapset.Set[Loc], pathOrdered []Loc, hBound int, vBound int, accCost int, bestCost *int) {
	if point.next.h < 0 || point.next.h >= mapp.hLength || point.next.v < 0 || point.next.v >= mapp.vLength {
		return
	}
	path.Add(point.next)
	pathOrdered = append(pathOrdered, point.next)
	accCost += mapp.data[point.next.v][point.next.h]
	if point.next.h > hBound {
		hBound = point.next.h
	}
	if point.next.v > vBound {
		vBound = point.next.v
	}
	if accCost > *bestCost {
		return
	}
	if point.next == mapp.endLoc {
		if accCost < *bestCost {
			*bestCost = accCost
			fmt.Printf("Cost: %+v\n", *bestCost)
			printMap(mapp, path)
		}
		return
	}
	nextPoints := []Point{point.moveInDirection("straight"), point.moveInDirection("right"), point.moveInDirection("left")}
	for _, nPoint := range nextPoints {
		if nPoint.next.h < hBound && nPoint.next.v < vBound {
			continue
		}
		if path.Contains(nPoint.next) {
			continue
		}
		pathLen := len(pathOrdered)
		if pathLen == 3 {
			previousLocs := pathOrdered[pathLen-3 : pathLen]
			validateLocs := append(previousLocs, nPoint.next)
			if allHValuesSame(validateLocs) || allVValuesSame(validateLocs) {
				continue
			}
		} else if pathLen >= 4 {
			previousLocs := pathOrdered[pathLen-4 : pathLen]
			validateLocs := append(previousLocs, nPoint.next)
			if allHValuesSame(validateLocs) || allVValuesSame(validateLocs) {
				continue
			}
		}
		newPath := path.Clone()
		nPoint.findWay(mapp, newPath, pathOrdered, hBound, vBound, accCost, bestCost)
	}
	return
}

func (loc *Loc) right() Loc {
	return Loc{loc.h + 1, loc.v}
}

type Map struct {
	data    [][]int
	hLength int
	vLength int
	endLoc  Loc
}

func (point *Point) moveInDirection(direction string) Point {
	hDelta := point.next.h - point.current.h
	vDelta := point.next.v - point.current.v
	nextPoint := Point{point.next, Loc{point.next.h + hDelta, point.next.v + vDelta}}
	switch direction {
	case "straight":
		return nextPoint.rotateCount(0)
	case "right":
		return nextPoint.rotateCount(1)
	case "left":
		return nextPoint.rotateCount(3)
	default:
		return nextPoint
	}
}

func (point *Point) rotateCount(count int) Point {
	for i := 0; i < count; i++ {
		*point = point.rotate()
	}
	return *point
}

func (point *Point) rotate() Point {
	// Calculate the horizontal and vertical distances
	hDist := point.next.h - point.current.h
	vDist := point.next.v - point.current.v
	// Rotate the coordinates:
	// - Swap the horizontal distance and vertical distance
	// - Negate the new horizontal distance (for clockwise rotation)
	rotatedNext := Loc{h: point.current.h - vDist, v: point.current.v + hDist}

	return Point{current: point.current, next: rotatedNext}
}

func allHValuesSame(locs []Loc) bool {
	// Get the first h value as a baseline
	firstH := locs[0].h

	// Iterate through the rest of the slice, comparing h values
	for _, loc := range locs[1:] {
		if loc.h != firstH {
			return false // If any h value differs, return false
		}
	}

	return true // If all h values match, return true
}

func allVValuesSame(locs []Loc) bool {
	firstV := locs[0].v // Get the first v value as a reference

	for _, loc := range locs[1:] { // Iterate through the rest of the slice
		if loc.v != firstV {
			return false // If any v value differs, return false
		}
	}

	return true // If all v values match, return true
}

func printMap(mapp *Map, path mapset.Set[Loc]) {
	for v := 0; v < mapp.vLength; v++ {
		for h := 0; h < mapp.hLength; h++ {
			loc := Loc{h, v}
			value := mapp.data[loc.v][loc.h]
			if path.Contains(loc) {
				fmt.Printf("%d", value)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

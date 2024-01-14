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
	starter.findWay(&mapp, path, 0, &bestCost)

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

func (point *Point) findWay(mapp *Map, path mapset.Set[Loc], accCost int, bestCost *int) {
	if point.next.h < 0 || point.next.h >= mapp.hLength || point.next.v < 0 || point.next.v >= mapp.vLength {
		return
	}
	path.Add(point.next)
	accCost += mapp.data[point.next.v][point.next.h]
	if accCost > *bestCost {
		return
	}
	if point.next == mapp.endLoc {
		if accCost < *bestCost {
			*bestCost = accCost
			fmt.Printf("Step: %+v\n", path.ToSlice())
		}
		return
	}
	nextPoints := []Point{point.moveInDirection("right"), point.moveInDirection("left"), point.moveInDirection("straight")}
	for _, nPoint := range nextPoints {
		if path.Contains(nPoint.next) {
			return
		}
		nPoint.findWay(mapp, path, accCost, bestCost)
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

func (loc *Loc) move(direction string) Loc {
	switch direction {
	case "right":
		return Loc{loc.h + 1, loc.v}
	case "left":
		return Loc{loc.h - 1, loc.v}
	case "up":
		return Loc{loc.h, loc.v - 1}
	case "down":
		return Loc{loc.h, loc.v + 1}
	default:
		return *loc
	}
}

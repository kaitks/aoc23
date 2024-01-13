package day16

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gammazero/deque"
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
	var mappData [][]string
	for _, rowStr := range strings.Split(data, "\n") {
		row := strings.Split(rowStr, "")
		mappData = append(mappData, row)
	}
	mapp := Map{mappData, len(mappData[0]), len(mappData)}
	starterLocs := []Loc{}
	for i := 0; i < mapp.hLength; i++ {
		starterLocs = append(starterLocs, Loc{i, -1}, Loc{i, mapp.vLength})
	}
	for i := 0; i < mapp.vLength; i++ {
		starterLocs = append(starterLocs, Loc{-1, i}, Loc{mapp.hLength, i})
	}
	maxSeen := 0
	for _, starter := range starterLocs {
		seen := calculateEnergizedTiles(starter, &mapp)
		if seen > maxSeen {
			maxSeen = seen
		}
	}

	fmt.Printf("Total: %+v\n", maxSeen)
	return maxSeen
}

func calculateEnergizedTiles(starter Loc, mapp *Map) int {
	starterNextLoc := starter
	if starter.v < 0 {
		starterNextLoc.v = 0
	} else if starter.v >= mapp.vLength {
		starterNextLoc.v = mapp.vLength - 1
	} else if starter.h < 0 {
		starterNextLoc.h = 0
	} else if starter.h >= mapp.hLength {
		starterNextLoc.h = mapp.hLength - 1
	}
	starterPoint := Point{starter, starterNextLoc}
	queue := deque.New[Point]()
	queue.PushBack(starterPoint)
	seen := mapset.NewSet[Point]()
	for {
		if queue.Len() == 0 {
			break
		}
		point := queue.PopBack()
		seen.Add(point)
		nextPoints := point.move(mapp)
		nextPoints = lo.Filter(nextPoints, func(point Point, _ int) bool {
			return !seen.Contains(point)
		})
		for _, nextPoint := range nextPoints {
			queue.PushBack(nextPoint)
		}
	}

	seenLoc := mapset.NewSet[Loc]()
	for _, point := range seen.ToSlice() {
		seenLoc.Add(point.next)
	}
	return seenLoc.Cardinality()
}

type Point struct {
	current Loc
	next    Loc
}

type Loc struct {
	h int
	v int
}

type Map struct {
	data    [][]string
	hLength int
	vLength int
}

func (point *Point) move(mapp *Map) []Point {
	next := point.next
	nextValue := mapp.data[next.v][next.h]
	//fmt.Printf("NextValue: %+v\n", nextValue)
	direction := point.direction()
	var nextPoints []Point
	switch nextValue {
	case "/":
		if direction == "right" {
			nextPoints = []Point{point.moveInDirection("up")}
		} else if direction == "left" {
			nextPoints = []Point{point.moveInDirection("down")}
		} else if direction == "up" {
			nextPoints = []Point{point.moveInDirection("right")}
		} else {
			nextPoints = []Point{point.moveInDirection("left")}
		}
	case "\\":
		if direction == "right" {
			nextPoints = []Point{point.moveInDirection("down")}
		} else if direction == "left" {
			nextPoints = []Point{point.moveInDirection("up")}
		} else if direction == "up" {
			nextPoints = []Point{point.moveInDirection("left")}
		} else {
			nextPoints = []Point{point.moveInDirection("right")}
		}
	case "|":
		if direction == "right" || direction == "left" {
			nextPoints = []Point{point.moveInDirection("up"), point.moveInDirection("down")}
		} else if direction == "up" || direction == "down" {
			nextPoints = []Point{point.moveInDirection(direction)}
		}
	case "-":
		if direction == "right" || direction == "left" {
			nextPoints = []Point{point.moveInDirection(direction)}
		} else if direction == "up" || direction == "down" {
			nextPoints = []Point{point.moveInDirection("left"), point.moveInDirection("right")}
		}
	case ".":
		nextPoints = []Point{point.moveInDirection(direction)}
	}
	nextPoints = lo.Filter(nextPoints, func(point Point, index int) bool {
		next := point.next
		return next.h >= 0 && next.h < mapp.hLength && next.v >= 0 && next.v < mapp.vLength
	})
	return nextPoints
}

func (point *Point) direction() string {
	current := point.current
	next := point.next
	direction := ""
	if next.h-current.h > 0 {
		direction = "right"
	} else if next.h-current.h < 0 {
		direction = "left"
	}
	if next.v-current.v > 0 {
		direction = "down"
	} else if next.v-current.v < 0 {
		direction = "up"
	}
	return direction
}

func (point *Point) moveInDirection(direction string) Point {
	next := point.next
	return Point{next, next.move(direction)}
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

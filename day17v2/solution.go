package day17v2

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"math"
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
	shortest := dj(&mapp)

	fmt.Printf("Total: %+v\n", shortest)
	return shortest
}

type Point struct {
	current Loc
	next    Loc
}

type Loc struct {
	h int
	v int
}

func dj(mapp *Map) int {
	unseen := mapset.NewSet[Loc]()
	score := map[Loc]Path{}
	for v := 0; v < mapp.vLength; v++ {
		for h := 0; h < mapp.hLength; h++ {
			loc := Loc{h, v}
			unseen.Add(loc)
			score[loc] = Path{[]Loc{}, math.MaxInt}
		}
	}
	score[Loc{0, 0}] = Path{[]Loc{}, 0}
	for {
		loc, found := findLowest(score, unseen)
		if !found {
			break
		} else {
			unseen.Remove(loc)
		}
		adjacentLocs := []Loc{loc.move("right"), loc.move("left"), loc.move("up"), loc.move("down")}
		adjacentLocs = lo.Filter(adjacentLocs, func(adj Loc, _ int) bool {
			return adj.h >= 0 && adj.h < mapp.hLength && adj.v >= 0 && adj.v < mapp.vLength
		})
		for _, adj := range adjacentLocs {
			currentPath, _ := score[loc]
			if lo.Contains(currentPath.path, adj) {
				continue
			}
			adjPath, _ := score[adj]
			distance := mapp.data[adj.v][adj.h]
			newDistance := currentPath.distance + distance
			if newDistance < adjPath.distance {
				adjPath.distance = newDistance
				adjPath.path = append(currentPath.path, adj)
				score[adj] = adjPath
			}
		}
	}
	bestPath := score[mapp.endLoc]
	fmt.Printf("Distance: %+v\n", bestPath.distance)
	printMap(mapp, mapset.NewSet[Loc](bestPath.path...))
	return bestPath.distance
}

func findLowest(score map[Loc]Path, unseen mapset.Set[Loc]) (Loc, bool) {
	minn := math.MaxInt
	found := false
	var loc Loc
	for k, path := range score {
		if path.distance < minn && unseen.Contains(k) {
			minn = path.distance
			loc = k
			found = true
		}
	}
	return loc, found
}

//
//func (point *Point) findWay(mapp *Map, path mapset.Set[Loc], pathOrdered []Loc, hBound int, vBound int, accCost int, bestCost *int) {
//	if point.next.h < 0 || point.next.h >= mapp.hLength || point.next.v < 0 || point.next.v >= mapp.vLength {
//		return
//	}
//	path.Add(point.next)
//	pathOrdered = append(pathOrdered, point.next)
//	accCost += mapp.data[point.next.v][point.next.h]
//	if point.next.h > hBound {
//		hBound = point.next.h
//	}
//	if point.next.v > vBound {
//		vBound = point.next.v
//	}
//	if accCost > *bestCost {
//		return
//	}
//	if point.next == mapp.endLoc {
//		if accCost < *bestCost {
//			*bestCost = accCost
//			fmt.Printf("Cost: %+v\n", *bestCost)
//			printMap(mapp, path)
//		}
//		return
//	}
//	nextPoints := []Point{point.moveInDirection("straight"), point.moveInDirection("right"), point.moveInDirection("left")}
//	for _, nPoint := range nextPoints {
//		if nPoint.next.h < hBound && nPoint.next.v < vBound {
//			continue
//		}
//		if path.Contains(nPoint.next) {
//			continue
//		}
//		pathLen := len(pathOrdered)
//		if pathLen == 3 {
//			previousLocs := pathOrdered[pathLen-3 : pathLen]
//			validateLocs := append(previousLocs, nPoint.next)
//			if allHValuesSame(validateLocs) || allVValuesSame(validateLocs) {
//				continue
//			}
//		} else if pathLen >= 4 {
//			previousLocs := pathOrdered[pathLen-4 : pathLen]
//			validateLocs := append(previousLocs, nPoint.next)
//			if allHValuesSame(validateLocs) || allVValuesSame(validateLocs) {
//				continue
//			}
//		}
//		newPath := path.Clone()
//		nPoint.findWay(mapp, newPath, pathOrdered, hBound, vBound, accCost, bestCost)
//	}
//	return
//}

func (loc *Loc) right() Loc {
	return Loc{loc.h + 1, loc.v}
}

type Map struct {
	data    [][]int
	hLength int
	vLength int
	endLoc  Loc
}

type Path struct {
	path     []Loc
	distance int
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

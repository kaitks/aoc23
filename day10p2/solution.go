package day10p2

import (
	"bufio"
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/samber/lo"
	"log"
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

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var tiles [][]string

	var sTile Tile
	vLength := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		tiles = append(tiles, row)
		for h, str := range row {
			if str == "S" {
				sTile = Tile{h, vLength, str}
			}
		}
		vLength++
	}

	hLength := len(tiles[0])
	mapp := Map{tiles, hLength, vLength}

	fmt.Printf("sTile: %+vLength\n", sTile)

	sLoc := sTile.toLoc()

	sNextIndex := ""

	sNexts := lo.FilterMap([]Loc{sLoc.up(), sLoc.down(), sLoc.right(), sLoc.left()}, func(loc Loc, index int) (Tile, bool) {
		Value, exists := mapp.get(loc.H, loc.V)
		if exists {
			sNextIndex += string(rune(index))
		}
		return Tile{loc.H, loc.V, Value}, exists
	})

	sRealValue := "."
	switch sNextIndex {
	case "01":
		sRealValue = "|"
	case "02":
		sRealValue = "L"
	case "03":
		sRealValue = "J"
	case "12":
		sRealValue = "F"
	case "13":
		sRealValue = "7"
	case "23":
		sRealValue = "-"
	}

	moves := mapset.NewSet[Loc]()

	for _, sNext := range sNexts {
		point := Point{sTile, sNext}
		thisMoves := mapset.NewSet[Loc]()
		canMove := true
		for canMove {
			thisMoves.Add(point.next.toLoc())
			nextPoint, canMovee := point.move(&mapp)
			point = nextPoint
			canMove = canMovee
		}
		if point.next == sTile {
			moves = thisMoves
			break
		}
	}

	LegitCross := mapset.NewSet("|", "L", "J")
	squares := mapset.NewSet[Loc]()

	for v := 0; v < vLength; v++ {
		inversions := 0
		for h := 0; h < hLength; h++ {
			loc := Loc{h, v}
			if moves.Contains(loc) {
				Value, _ := mapp.get(loc.H, loc.V)
				if LegitCross.Contains(Value) || (Value == "S" && LegitCross.Contains(sRealValue)) {
					inversions++
				}
			} else {
				if inversions%2 != 0 {
					squares.Add(loc)
				}
			}
		}
	}

	//fmt.Printf("Move: %+v\n", moves.ToSlice())
	//fmt.Printf("Square: %+v\n", squares.ToSlice())
	fmt.Printf("Total: %+v\n", squares.Cardinality())

	return squares.Cardinality()
}

type Tile struct {
	H     int
	V     int
	Value string
}

type Loc struct {
	H int
	V int
}

type Point struct {
	current Tile
	next    Tile
}

type Map struct {
	Data    [][]string
	hLength int
	vLength int
}

func (tile *Tile) toLoc() Loc {
	return Loc{tile.H, tile.V}
}

func (loc *Loc) right() Loc {
	return Loc{loc.H + 1, loc.V}
}

func (loc *Loc) left() Loc {
	return Loc{loc.H - 1, loc.V}
}

func (loc *Loc) up() Loc {
	return Loc{loc.H, loc.V - 1}
}

func (loc *Loc) down() Loc {
	return Loc{loc.H, loc.V + 1}
}

func (mapp *Map) get(h int, v int) (string, bool) {
	if h < 0 || h >= mapp.hLength || v < 0 || v >= mapp.vLength {
		return "", false
	} else {
		return mapp.Data[v][h], true
	}
}

func (point *Point) move(mapp *Map) (Point, bool) {
	nextt, canMove := point.current.move(point.next, mapp)
	if canMove {
		return Point{point.next, nextt}, true
	} else {
		return *point, false
	}
}

func (tile *Tile) move(next Tile, mapp *Map) (Tile, bool) {
	var nextt Tile
	nexttLoc := Loc{-1, -1}
	nextLoc := next.toLoc()
	hDelta := next.H - tile.H
	vDelta := next.V - tile.V

	switch next.Value {
	case "|":
		if hDelta == 0 {
			nexttLoc = Loc{nextLoc.H, nextLoc.V + vDelta}
		}
	case "-":
		if vDelta == 0 {
			nexttLoc = Loc{nextLoc.H + hDelta, nextLoc.V}
		}
	case "L":
		if vDelta == 1 {
			nexttLoc = nextLoc.right()
		} else if hDelta == -1 {
			nexttLoc = nextLoc.up()
		}
	case "J":
		if vDelta == 1 {
			nexttLoc = nextLoc.left()
		} else if hDelta == 1 {
			nexttLoc = nextLoc.up()
		}
	case "7":
		if hDelta == 1 {
			nexttLoc = nextLoc.down()
		} else if vDelta == -1 {
			nexttLoc = nextLoc.left()
		}
	case "F":
		if hDelta == -1 {
			nexttLoc = nextLoc.down()
		} else if vDelta == -1 {
			nexttLoc = nextLoc.right()
		}
	}
	nextValue, exists := (*mapp).get(nexttLoc.H, nexttLoc.V)
	if exists {
		nextt = Tile{nexttLoc.H, nexttLoc.V, nextValue}
		return nextt, true
	} else {
		return next, false
	}
}

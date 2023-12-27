package day10

import (
	"bufio"
	"fmt"
	"github.com/samber/lo"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func part1(fileName string) int {
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
	v := 0

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		tiles = append(tiles, row)
		for h, str := range row {
			if str == "S" {
				sTile = Tile{h, v, str}
			}
		}
		v++
	}

	hLength := len(tiles[0])
	mapp := Map{tiles, hLength, v}

	fmt.Printf("sTile: %+v\n", sTile)

	sLoc := sTile.toLoc()

	sNexts := lo.FilterMap([]Loc{sLoc.up(), sLoc.down(), sLoc.right(), sLoc.left()}, func(loc Loc, _ int) (Tile, bool) {
		Value, exists := mapp.get(loc.H, loc.V)
		return Tile{loc.H, loc.V, Value}, exists
	})

	var moves []string

	for _, sNext := range sNexts {
		point := Point{sTile, sNext}
		var thisMoves []string
		canMove := true
		for canMove {
			thisMoves = append(thisMoves, point.next.Value)
			nextPoint, canMovee := point.move(&mapp)
			point = nextPoint
			canMove = canMovee
		}
		if point.next.Value == "S" {
			moves = thisMoves
			break
		}
	}

	fmt.Printf("Move: %+v\n", moves)
	fmt.Printf("Total: %+v\n", len(moves)/2)

	return len(moves) / 2
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

func right(h int, v int) (int, int) {
	return h + 1, v
}

func left(h int, v int) (int, int) {
	return h - 1, v
}

func up(h int, v int) (int, int) {
	return h, v - 1
}

func down(h int, v int) (int, int) {
	return h, v + 1
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
	nextH := -1
	nextV := -1
	hDelta := next.H - tile.H
	vDelta := next.V - tile.V

	switch next.Value {
	case "|":
		if hDelta == 0 {
			nextH = next.H
			nextV = next.V + vDelta
		}
	case "-":
		if vDelta == 0 {
			nextH = next.H + hDelta
			nextV = next.V
		}
	case "L":
		if vDelta == 1 {
			nextH, nextV = right(next.H, next.V)
		} else if hDelta == -1 {
			nextH, nextV = up(next.H, next.V)
		}
	case "J":
		if vDelta == 1 {
			nextH, nextV = left(next.H, next.V)
		} else if hDelta == 1 {
			nextH, nextV = up(next.H, next.V)
		}
	case "7":
		if hDelta == 1 {
			nextH, nextV = down(next.H, next.V)
		} else if vDelta == -1 {
			nextH, nextV = left(next.H, next.V)
		}
	case "F":
		if hDelta == -1 {
			nextH, nextV = down(next.H, next.V)
		} else if vDelta == -1 {
			nextH, nextV = right(next.H, next.V)
		}
	}
	if nextH >= 0 && nextV >= 0 {
		nextValue, exists := (*mapp).get(nextH, nextV)
		if exists {
			nextt = Tile{nextH, nextV, nextValue}
			return nextt, true
		} else {
			return next, false
		}
	} else {
		return next, false
	}
}

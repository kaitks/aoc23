package day10

import (
	"bufio"
	"fmt"
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

	acc := 0

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

	sNext := Tile{sTile.H + 1, sTile.V, tiles[sTile.V][sTile.H+1]}

	point := Point{sTile, sNext}
	moves := []string{}
	canMove := true
	for point.next.Value != "S" && canMove {
		nextPoint, canMovee := point.move(&tiles)
		moves = append(moves, nextPoint.next.Value)
		point = nextPoint
		canMove = canMovee
	}

	fmt.Printf("Move: %+v\n", moves)
	fmt.Printf("Total: %+v\n", tiles)

	return acc
}

type Tile struct {
	H     int
	V     int
	Value string
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
	if h < 0 || h > mapp.hLength || v < 0 || v > mapp.vLength {
		return "", false
	} else {
		return mapp.Data[v][h], true
	}
}

func (point *Point) move(mapp *[][]string) (Point, bool) {
	nextt, canMove := point.current.move(point.next, mapp)
	return Point{point.next, nextt}, canMove
}

func (tile *Tile) move(next Tile, mapp *[][]string) (Tile, bool) {
	var nextt Tile
	var nextH int
	var nextV int
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
			nextH, nextV = right(next.H, nextV)
		} else if hDelta == -1 {
			nextH, nextV = up(next.H, nextV)
		}
	case "J":
		if vDelta == 1 {
			nextH, nextV = left(next.H, nextV)
		} else if hDelta == 1 {
			nextH, nextV = up(next.H, nextV)
		}
	case "7":
		if hDelta == 1 {
			nextH, nextV = down(next.H, nextV)
		} else if vDelta == -1 {
			nextH, nextV = left(next.H, nextV)
		}
	case "F":
		if hDelta == -1 {
			nextH, nextV = down(next.H, nextV)
		} else if vDelta == -1 {
			nextH, nextV = right(next.H, nextV)
		}
	}
	canMove := false
	if &nextH != nil && &nextV != nil {
		canMove = true
		nextValue := (*mapp)[nextV][nextH]
		nextt = Tile{nextH, nextV, nextValue}
		return nextt, canMove
	} else {
		return next, canMove
	}
}

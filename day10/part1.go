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
	loc  Tile
	next Tile
}

func (point *Point) move(mapp *[][]string) (Point, bool) {
	nextt, canMove := point.loc.move(point.next, mapp)
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
			nextH = next.H + 1
			nextV = next.V
		} else if hDelta == -1 {
			nextH = next.H
			nextV = next.V - 1
		}
	case "J":
		if vDelta == 1 {
			nextH = next.H - 1
			nextV = next.V
		} else if hDelta == 1 {
			nextH = next.H
			nextV = next.V - 1
		}
	case "7":
		if hDelta == 1 {
			nextH = next.H
			nextV = next.V + 1
		} else if vDelta == -1 {
			nextH = next.H - 1
			nextV = next.V
		}
	case "F":
		if hDelta == -1 {
			nextH = next.H
			nextV = next.V + 1
		} else if vDelta == -1 {
			nextH = next.H + 1
			nextV = next.V
		}
	}
	canMove := false
	if &nextH != nil && &nextV != nil {
		canMove = true
		nextValue := (*mapp)[nextV][nextH]
		nextt = Tile{nextH, nextV, nextValue}
	}
	return nextt, canMove
}

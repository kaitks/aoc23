package day20

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gammazero/deque"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"strings"
)

func solution(fileName string, maxStep int) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	// Create a scanner to read the file line by line
	raw, _ := os.ReadFile(filePath)
	data := string(raw)
	mapp := parseMap(data)
	stepQueue := deque.Deque[PosStep]{}
	seen := mapset.NewSet[Pos]()
	stepQueue.PushBack(PosStep{Pos: mapp.S, Step: 0})

	for {
		if stepQueue.Len() == 0 {
			break
		}
		pos := stepQueue.PopFront()
		if pos.Step == maxStep {
			seen.Add(pos.Pos)
			continue
		}
		nextPositions := pos.step()
		nextPositions = lo.Filter(nextPositions, func(next Pos, _ int) bool {
			return next.X >= 0 && next.X < mapp.Width && next.Y >= 0 && next.Y < mapp.Height && !mapp.Rocks.Contains(next)
		})
		for _, next := range nextPositions {
			stepQueue.PushBack(PosStep{Pos: next, Step: pos.Step + 1})
		}
	}

	total := seen.Cardinality()
	fmt.Printf("\nTotal: %+v\n", total)
	return total
}

type Map struct {
	Grid   [][]int32
	Width  int
	Height int
	Rocks  mapset.Set[Pos]
	S      Pos
}

type Pos struct {
	X int
	Y int
}

type PosStep struct {
	Pos
	Step int
}

func (pos *Pos) step() []Pos {
	x := pos.X
	y := pos.Y
	return []Pos{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}
}

func parseMap(data string) *Map {
	rows := strings.Split(data, "\n")
	grids := [][]int32{}
	var S Pos
	Rocks := mapset.NewSet[Pos]()
	for y, rowStr := range rows {
		row := []int32{}
		for x, char := range rowStr {
			row = append(row, char)
			if char == 'S' {
				S = Pos{x, y}
			} else if char == '#' {
				Rocks.Add(Pos{x, y})
			}
		}
		grids = append(grids, row)
	}
	Height := len(rows)
	Width := len(grids[0])
	return &Map{Grid: grids, Width: Width, Height: Height, Rocks: Rocks, S: S}
}

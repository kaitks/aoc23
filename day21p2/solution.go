package day21p2

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
	var a []int
	n := 0
	for {
		a = append(a, getReachablePosQuadraticByN(mapp, maxStep, n))
		n++
		if len(a) >= 4 {
			fd := []int{a[1] - a[0], a[2] - a[1], a[3] - a[2]}
			sd := []int{fd[1] - fd[0], fd[2] - fd[1]}
			fmt.Printf("fd: %v\n", fd)
			fmt.Printf("sd: %v\n\n", sd)
			if sd[1] == sd[0] {
				break
			} else {
				a = a[1:]
			}
		}
	}

	alpha := a[0]
	beta := a[1]
	gamma := a[2]

	reach := getReachablePosQuadratic(alpha, beta, gamma, mapp, maxStep, n)
	fmt.Printf("\nStep: %d, Total: %+v\n", maxStep, reach)
	return reach
}

func getReachablePosQuadratic(alpha, beta, gamma int, mapp *Map, maxStep int, n int) int {
	c := alpha
	a := (gamma - 2*beta + c) / 2
	b := beta - c - a
	size := mapp.Height
	offset := n - 4
	step := maxStep/(2*size) - offset
	return a*step*step + b*step + c
}

func getReachablePosQuadraticByN(mapp *Map, maxStep int, n int) int {
	size := mapp.Height
	original := maxStep % (2 * size)
	step := original + 2*n*size
	return getReachablePos(mapp, step)
}

func getReachablePos(mapp *Map, maxStep int) int {
	stepQueue := deque.Deque[PosStep]{}
	seenMap := map[Pos]int{}
	stepQueue.PushBack(PosStep{Pos: mapp.S, Step: 0})

	remain := maxStep % 2

	for {
		if stepQueue.Len() == 0 {
			break
		}
		pos := stepQueue.PopFront()
		posStepRemain := pos.Step % 2
		if posStepRemain == remain {
			seenStep, exists := seenMap[pos.Pos]
			if !exists {
				seenMap[pos.Pos] = pos.Step
			} else {
				if pos.Step < seenStep {
					seenMap[pos.Pos] = pos.Step
				} else {
					continue
				}
			}
		}
		if pos.Step == maxStep {
			continue
		}
		nextPositions := pos.step()
		nextPositions = lo.Filter(nextPositions, func(next Pos, _ int) bool {
			var relativeX int
			var relativeY int
			if next.X < 0 {
				relativeX = mapp.Width + (next.X % mapp.Width)
			} else {
				relativeX = next.X % mapp.Width
			}
			if next.Y < 0 {
				relativeY = mapp.Height + (next.Y % mapp.Height)
			} else {
				relativeY = next.Y % mapp.Height
			}
			return !mapp.Rocks.Contains(Pos{relativeX, relativeY})
		})
		for _, next := range nextPositions {
			stepQueue.PushBack(PosStep{Pos: next, Step: pos.Step + 1})
		}
	}

	total := len(seenMap)
	//fmt.Printf("\nStep: %d, Total: %+v\n", maxStep, total)
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
	var grids [][]int32
	var S Pos
	Rocks := mapset.NewSet[Pos]()
	for y, rowStr := range rows {
		var row []int32
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

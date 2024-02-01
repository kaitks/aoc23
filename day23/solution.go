package day23

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
	mapp := parseData(data)
	start := Pos{1, 0}
	end := Pos{mapp.width - 2, mapp.height - 1}
	queue := deque.New[Point]()
	history := mapset.NewSet[Pos]()
	history.Add(start)
	queue.PushBack(Point{start, &history})
	var possiblePath []*mapset.Set[Pos]
	for {
		if queue.Len() == 0 {
			break
		}
		point := queue.PopFront()
		current := point.current
		var nextPos []Pos
		switch mapp.grids[current.y][current.x] {
		case '.':
			nextPos = append(nextPos, []Pos{{current.x - 1, current.y}, {current.x + 1, current.y}, {current.x, current.y - 1}, {current.x, current.y + 1}}...)
		case '^':
			nextPos = append(nextPos, Pos{current.x, current.y - 1})
		case '>':
			nextPos = append(nextPos, Pos{current.x + 1, current.y})
		case 'v':
			nextPos = append(nextPos, Pos{current.x, current.y + 1})
		case '<':
			nextPos = append(nextPos, Pos{current.x - 1, current.y})
		}
		nextPos = lo.Filter(nextPos, func(pos Pos, _ int) bool {
			return pos.x >= 0 && pos.x < mapp.width && pos.y >= 0 && pos.y < mapp.height && mapp.grids[pos.y][pos.x] != '#' && !(*(point.history)).Contains(pos)
		})
		for _, pos := range nextPos {
			nextHistory := (*(point.history)).Clone()
			nextHistory.Add(pos)
			nextPoint := Point{pos, &nextHistory}
			if pos == end {
				possiblePath = append(possiblePath, nextPoint.history)
			} else {
				queue.PushBack(nextPoint)
			}
		}
	}
	total := (*(lo.MaxBy(possiblePath, func(a, b *mapset.Set[Pos]) bool {
		return (*a).Cardinality() > (*b).Cardinality()
	}))).Cardinality() - 1
	fmt.Printf("Longest Hike: %+v\n", total)
	return total
}

type Map struct {
	grids         [][]int32
	width, height int
}

type Point struct {
	current Pos
	history *mapset.Set[Pos]
}

func parseData(data string) *Map {
	rows := strings.Split(data, "\n")
	grids := [][]int32{}
	for _, rowStr := range rows {
		row := []int32{}
		for _, char := range rowStr {
			row = append(row, char)
		}
		grids = append(grids, row)
	}
	return &Map{grids, len(grids[0]), len(grids)}
}

type Pos struct {
	x, y int
}

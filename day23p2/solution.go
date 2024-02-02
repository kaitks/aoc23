package day23p2

import (
	"fmt"
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
	history := History{}
	history.Add(start)
	queue.PushBack(Point{start, &history})
	var possiblePath []*History
	for {
		if queue.Len() == 0 {
			break
		}
		point := queue.PopFront()
		current := point.current
		var nextPos []Pos
		nextPos = append(nextPos, []Pos{{current.x - 1, current.y}, {current.x + 1, current.y}, {current.x, current.y - 1}, {current.x, current.y + 1}}...)
		nextPos = lo.Filter(nextPos, func(pos Pos, _ int) bool {
			return pos.x >= 0 && pos.x < mapp.width && pos.y >= 0 && pos.y < mapp.height && mapp.grids[pos.y][pos.x] != '#' && !(*(point.history)).Contains(pos)
		})
		if len(nextPos) == 1 {
			pos := nextPos[0]
			nextHistory := point.history
			(*nextHistory).Add(pos)
			nextPoint := Point{pos, point.history}
			if pos == end {
				possiblePath = append(possiblePath, nextPoint.history)
			} else {
				queue.PushBack(nextPoint)
			}
		} else {
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
	}
	total := len(*(lo.MaxBy(possiblePath, func(a, b *History) bool {
		return len(*a) > len(*b)
	}))) - 1
	fmt.Printf("Longest Hike: %+v\n", total)
	return total
}

type Map struct {
	grids         [][]int32
	width, height int
}

type Point struct {
	current Pos
	history *History
}

func parseData(data string) *Map {
	rows := strings.Split(data, "\n")
	var grids [][]int32
	for _, rowStr := range rows {
		var row []int32
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

type History map[Pos]int

func (history *History) Add(pos Pos) {
	value, exists := (*history)[pos]
	if !exists {
		value = 0
	}
	(*history)[pos] = value + 1
}

func (history *History) Clone() History {
	clone := History{}
	for k, v := range *history {
		clone[k] = v
	}
	return clone
}

func (history *History) Contains(pos Pos) bool {
	value, exists := (*history)[pos]
	if !exists {
		value = 0
	}
	return value >= 2
}

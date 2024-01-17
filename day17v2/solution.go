package day17v2

import (
	"container/heap"
	"fmt"
	"github.com/samber/lo"
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
	var grid [][]int
	for _, rowStr := range strings.Split(data, "\n") {
		row := lo.Map(strings.Split(rowStr, ""), func(str string, _ int) int {
			value, _ := strconv.Atoi(str)
			return value
		})
		grid = append(grid, row)
	}
	hLength := len(grid[0])
	vLength := len(grid)
	mapp := Map{grid, hLength, vLength, Pos{hLength - 1, vLength - 1}}
	part1 := Dj(&mapp, 1, 3)
	part2 := Dj(&mapp, 4, 10)

	fmt.Printf("Part1: %+v\n", part1)
	fmt.Printf("Part2: %+v\n", part2)
	return part1
}

type Pos struct {
	x int
	y int
}

func Dj(mapp *Map, minSteps, maxSteps int) int {
	graph := NewGraph(mapp)
	vertices := graph.vertices
	vertices[0].direction = PlaneUndecided
	vertices[0].total = 0

	pq := make(PriorityQueue, len(vertices))
	for i := 0; i < len(vertices); i++ {
		vertices[i].index = i
		pq[i] = &vertices[i]
	}
	heap.Init(&pq)

	var u *Vertex

	for {
		u = heap.Pop(&pq).(*Vertex)
		if (*u).pos == mapp.endPos {
			break
		}
		edges := graph.getEdges(u, minSteps, maxSteps)
		for _, e := range edges {
			total := u.total + e.calculatedCost
			if e.total > total {
				e.total = total
				heap.Fix(&pq, e.index)
			}
		}
	}
	return u.total
}

type PriorityQueue []*Vertex

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].total < (*pq)[j].total
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index = i
	(*pq)[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	index := (*pq).Len()
	item := x.(*Vertex)
	item.index = index
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	n := (*pq).Len()
	item := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	item.index = -1
	return item
}

type Map struct {
	grid    [][]int
	hLength int
	vLength int
	endPos  Pos
}

type Graph struct {
	vertices []Vertex
	width    int
	height   int
}

func (graph *Graph) getEdges(u *Vertex, minSteps, maxSteps int) []*Vertex {
	var e []*Vertex
	if u.direction == PlaneVertical || u.direction == PlaneUndecided {
		for calculatedCost, i := 0, 1; i <= maxSteps; i++ {
			v := graph.getVertexByPos(u.pos.x, u.pos.y+i, PlaneHorizontal)
			if v != nil {
				calculatedCost += v.cost
				if i >= minSteps {
					v.calculatedCost = calculatedCost
					e = append(e, v)
				}
			}
		}
		for calculatedCost, i := 0, 1; i <= maxSteps; i++ {
			v := graph.getVertexByPos(u.pos.x, u.pos.y-i, PlaneHorizontal)
			if v != nil {
				calculatedCost += v.cost
				if i >= minSteps {
					v.calculatedCost = calculatedCost
					e = append(e, v)
				}
			}
		}
	}
	if u.direction == PlaneHorizontal || u.direction == PlaneUndecided {
		for calculatedCost, i := 0, 1; i <= maxSteps; i++ {
			v := graph.getVertexByPos(u.pos.x+i, u.pos.y, PlaneVertical)
			if v != nil {
				calculatedCost += v.cost
				if i >= minSteps {
					v.calculatedCost = calculatedCost
					e = append(e, v)
				}
			}
		}
		for calculatedCost, i := 0, 1; i <= maxSteps; i++ {
			v := graph.getVertexByPos(u.pos.x-i, u.pos.y, PlaneVertical)
			if v != nil {
				calculatedCost += v.cost
				if i >= minSteps {
					v.calculatedCost = calculatedCost
					e = append(e, v)
				}
			}
		}
	}
	return e
}

func (graph *Graph) getVertexByPos(x int, y int, direction int) *Vertex {
	if x < 0 || y < 0 || y >= graph.height || x >= graph.width {
		return nil
	}
	return &graph.vertices[y*graph.width*2+x*2+direction]
}

type Vertex struct {
	index          int
	pos            Pos
	direction      int
	cost           int
	calculatedCost int
	total          int
}

func NewGraph(mapp *Map) Graph {
	var vertices = make([]Vertex, 0, mapp.hLength*mapp.vLength*2)
	for y, row := range mapp.grid {
		for x, cost := range row {
			vertices = append(vertices, Vertex{pos: Pos{x, y}, direction: PlaneVertical, cost: cost, total: 1 << 30})
			vertices = append(vertices, Vertex{pos: Pos{x, y}, direction: PlaneHorizontal, cost: cost, total: 1 << 30})
		}
	}
	return Graph{vertices, mapp.hLength, mapp.vLength}
}

const (
	PlaneVertical = iota
	PlaneHorizontal
	PlaneUndecided // special plane for start position
)

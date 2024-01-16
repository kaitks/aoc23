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
	shortest := Dj(&mapp)

	fmt.Printf("Total: %+v\n", shortest)
	return shortest
}

type Pos struct {
	x int
	y int
}

func Dj(mapp *Map) int {
	graph := NewGraph(mapp)
	vertices := graph.vertices
	vertices[0].direction = PLANE_UNDECIDED
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
		edges := graph.getEdges(u)
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

func (graph *Graph) getEdges(u *Vertex) []*Vertex {
	var e []*Vertex
	if u.direction == PLANE_VERTICAL || u.direction == PLANE_UNDECIDED {
		for calculatedCost, i := 0, 1; i <= 3; i++ {
			v := graph.getVertexByPos(u.pos.x, u.pos.y+i, PLANE_HORIZONTAL)
			if v != nil {
				calculatedCost += v.cost
				v.calculatedCost = calculatedCost
				e = append(e, v)
			}
		}
		for calculatedCost, i := 0, 1; i <= 3; i++ {
			v := graph.getVertexByPos(u.pos.x, u.pos.y-i, PLANE_HORIZONTAL)
			if v != nil {
				calculatedCost += v.cost
				v.calculatedCost = calculatedCost
				e = append(e, v)
			}
		}
	}
	if u.direction == PLANE_HORIZONTAL || u.direction == PLANE_UNDECIDED {
		for calculatedCost, i := 0, 1; i <= 3; i++ {
			v := graph.getVertexByPos(u.pos.x+i, u.pos.y, PLANE_VERTICAL)
			if v != nil {
				calculatedCost += v.cost
				v.calculatedCost = calculatedCost
				e = append(e, v)
			}
		}
		for calculatedCost, i := 0, 1; i <= 3; i++ {
			v := graph.getVertexByPos(u.pos.x-i, u.pos.y, PLANE_VERTICAL)
			if v != nil {
				calculatedCost += v.cost
				v.calculatedCost = calculatedCost
				e = append(e, v)
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
			vertices = append(vertices, Vertex{pos: Pos{x, y}, direction: PLANE_VERTICAL, cost: cost, total: 1 << 30})
			vertices = append(vertices, Vertex{pos: Pos{x, y}, direction: PLANE_HORIZONTAL, cost: cost, total: 1 << 30})
		}
	}
	return Graph{vertices, mapp.hLength, mapp.vLength}
}

const (
	PLANE_VERTICAL = iota
	PLANE_HORIZONTAL
	PLANE_UNDECIDED // special plane for start position
)

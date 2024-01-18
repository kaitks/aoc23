package day18p2

import (
	"fmt"
	"math"
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
	verticesStr := strings.Split(data, "\n")
	currentPoint := Point{}
	vertices := make([]Point, 0, len(verticesStr))
	perimeter := 0
	for _, verticeStr := range verticesStr {
		instruction := strings.Fields(verticeStr)[2]
		iLen := len(instruction)
		directionInt, _ := strconv.Atoi(instruction[iLen-2 : iLen-1])
		directionMap := map[int]string{}
		directionMap[0] = "R"
		directionMap[1] = "D"
		directionMap[2] = "L"
		directionMap[3] = "U"
		direction := directionMap[directionInt]
		stepStr := instruction[2 : iLen-2]
		step64, _ := strconv.ParseInt(stepStr, 16, 32)
		step := int(step64)
		perimeter += step
		currentPoint = currentPoint.move(direction, step)
		vertices = append(vertices, currentPoint)
	}
	area := polygonArea(vertices)
	total := area + perimeter/2 + 1
	fmt.Printf("Total: %+v\n", total)
	return total
}

type Point struct {
	X, Y int
}

func (point *Point) move(direction string, step int) Point {
	switch direction {
	case "R":
		return Point{point.X + step, point.Y}
	case "L":
		return Point{point.X - step, point.Y}
	case "U":
		return Point{point.X, point.Y - step}
	case "D":
		return Point{point.X, point.Y + step}
	default:
		return Point{}
	}
}

func polygonArea(vertices []Point) int {
	area := 0
	n := len(vertices)
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += vertices[i].X * vertices[j].Y
		area -= vertices[i].Y * vertices[j].X
	}
	area = int(math.Abs(float64(area))) / 2 // Convert to float for division, then back to int
	return area
}

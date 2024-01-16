package day17ref

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

func TestSolution(t *testing.T) {
	timeStart := time.Now()

	input, err := os.ReadFile("test_data")
	if err != nil {
		log.Fatal(err)
	}
	grid := parse(input)

	a1 := Dijkstra(grid, 1, 3)
	a2 := Dijkstra(grid, 4, 10)
	fmt.Printf("--- Day 17: Clumsy Crucible ---\n")
	fmt.Printf("Part 1: %d\n", a1)
	fmt.Printf("Part 2: %d\n", a2)
	fmt.Printf("Time: %.2fms\n", float64(time.Since(timeStart).Microseconds())/1000)
}

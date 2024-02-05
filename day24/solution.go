package day24

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func solution(fileName string, min, max float64) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	// Create a scanner to read the file line by line
	raw, _ := os.ReadFile(filePath)
	data := string(raw)
	stones := parseData(data)
	ssfs := lo.Map(stones, func(stone Stone, _ int) StoneStandardForm {
		x, y, vx, vy := stone.x, stone.y, stone.vx, stone.vy
		a := vy
		b := -vx
		c := vy*x - vx*y
		s := StoneStandardForm{a, b, c}
		return s
	})

	intersect := 0
	for i := 0; i < len(ssfs); i++ {
		stoneA := stones[i]
		sA := ssfs[i]
		for j := i + 1; j < len(ssfs); j++ {
			stoneB := stones[j]
			sB := ssfs[j]
			isIntersect, x, y := sA.intersect(&sB)
			intersectInTheFuture := isIntersect && (x-stoneA.x)/stoneA.vx > 0 && (x-stoneB.x)/stoneB.vx > 0
			intersectInBound := intersectInTheFuture && min <= x && x <= max && min <= y && y <= max
			if intersectInBound {
				//fmt.Printf("Stones: %+v, %v\n", stoneA, stoneB)
				intersect++
			}
		}
	}
	fmt.Printf("Stones: %+v\n", intersect)
	return intersect
}

type Stone struct {
	x, y, z, vx, vy, vz float64
}

type StoneStandardForm struct {
	a, b, c float64
}

func (s *StoneStandardForm) print() {
	fmt.Printf("Stone: %fx + %fy = %f\n", s.a, s.b, s.c)
}

func (s *StoneStandardForm) intersect(o *StoneStandardForm) (bool, float64, float64) {
	a1, b1, c1 := s.a, s.b, s.c
	a2, b2, c2 := o.a, o.b, o.c
	if a1*b2 == b1*a2 {
		return false, 0, 0
	}
	x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
	y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)
	return true, x, y
}

func parseData(data string) []Stone {
	rows := strings.Split(data, "\n")
	stones := make([]Stone, 0, len(rows))
	regex := regexp.MustCompile(`-?\d+`)
	for _, rowStr := range rows {
		m := lo.Map(regex.FindAllString(rowStr, -1), func(number string, _ int) float64 {
			n, _ := strconv.ParseFloat(number, 64)
			return n
		})
		stones = append(stones, Stone{m[0], m[1], m[2], m[3], m[4], m[5]})
	}
	return stones
}

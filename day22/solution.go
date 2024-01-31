package day22

import (
	"fmt"
	"github.com/samber/lo"
	"math"
	"os"
	"path/filepath"
	"sort"
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
	bricks := parseData(data)
	bricks = append(bricks, &Brick{id: 0, x: Range{0, math.MaxInt}, y: Range{0, math.MaxInt}, z: Range{0, 0}})
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].z.start < bricks[j].z.end
	})
	for i := 1; i < len(bricks); i++ {
		canDrop := true
		top := bricks[i]
		bottoms := make([]*Brick, 0, len(bricks[:i]))
		for _, brick := range bricks[:i] {
			if brick.z.end < top.z.start {
				bottoms = append(bottoms, brick)
			}
		}
		sort.Slice(bottoms, func(i, j int) bool {
			return bottoms[i].z.end > bottoms[j].z.end
		})
		for _, bottom := range bottoms {
			if top.overlap(bottom) {
				if canDrop {
					newZ := Range{bottom.z.end + 1, bottom.z.end + 1 + top.z.end - top.z.start}
					top.z = newZ
					canDrop = false
				}
				if top.z.start == bottom.z.end+1 {
					bottom.support = append(bottom.support, top)
					top.supportedBy = append(top.supportedBy, bottom)
				}
				if top.z.start > bottom.z.end+1 {
					break
				}
			}
		}
	}

	var disintegrate []*Brick

	for _, brick := range bricks {
		canBeRemoved := true
		for _, supporting := range brick.support {
			if len(supporting.supportedBy) == 1 {
				canBeRemoved = false
				break
			}
		}
		if canBeRemoved {
			//for _, supporting := range brick.support {
			//	supporting.supportedBy = slices.DeleteFunc(supporting.supportedBy, func(supportedBy *Brick) bool {
			//		return supportedBy == brick
			//	})
			//}
			disintegrate = append(disintegrate, brick)
		}
	}
	result := len(disintegrate)
	fmt.Printf("Bricks: %+v\n", result)
	return result
}

type Brick struct {
	id          int
	x, y, z     Range
	support     []*Brick
	supportedBy []*Brick
}

type Range struct {
	start, end int
}

func parseData(data string) []*Brick {
	rows := strings.Split(data, "\n")
	var bricks []*Brick
	for i, row := range rows {
		ranges := strings.Split(row, "~")
		start := lo.Map(strings.Split(ranges[0], ","), func(str string, _ int) int {
			number, _ := strconv.Atoi(str)
			return number
		})
		end := lo.Map(strings.Split(ranges[1], ","), func(str string, _ int) int {
			number, _ := strconv.Atoi(str)
			return number
		})
		bricks = append(bricks, &Brick{id: i + 1, x: Range{start[0], end[0]}, y: Range{start[1], end[1]}, z: Range{start[2], end[2]}})
	}
	return bricks
}

type Bricks []*Brick

func (brick *Brick) overlap(other *Brick) bool {
	return isRangeOverlap(brick.x, other.x) && isRangeOverlap(brick.y, other.y)
}

func isRangeOverlap(range1, range2 Range) bool {
	// Check if either range's start falls within the other's range
	if range1.start <= range2.end && range2.start <= range1.end {
		return true
	}
	// Check if either range encompasses the other entirely
	if range1.start <= range2.start && range1.end >= range2.end {
		return true
	}
	if range2.start <= range1.start && range2.end >= range1.end {
		return true
	}
	// No overlap
	return false
}

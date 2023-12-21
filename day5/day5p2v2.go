package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func day5p2v2(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	raw, _ := os.ReadFile(filePath)
	data := string(raw)

	sections := strings.Split(data, "\n\n")
	seedRaws := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	seeds := []Seed{}
	resources := [][]Resource{}

	for i := 0; i < len(seedRaws); i = i + 2 {
		seedStart, _ := strconv.Atoi(seedRaws[i])
		seedLength, _ := strconv.Atoi(seedRaws[i+1])
		seeds = append(seeds, Seed{seedStart, seedLength})
	}

	for i := 1; i < len(sections); i++ {
		mapData := strings.Split(strings.Split(sections[i], ":\n")[1], "\n")
		resourceGroup := []Resource{}
		for _, mapString := range mapData {
			fields := strings.Split(mapString, " ")
			destination, _ := strconv.Atoi(fields[0])
			source, _ := strconv.Atoi(fields[1])
			rangeLength, _ := strconv.Atoi(fields[2])
			resourceGroup = append(resourceGroup, mkResource(source, destination, rangeLength))
		}
		resources = append(resources, resourceGroup)
	}

	results := make(chan int)

	for _, seed := range seeds {
		go func(seed Seed, c chan int) {
			result := getDestinationBySeedV2(seed, resources)
			c <- result
		}(seed, results)
	}

	locs := []int{}

	for i := 0; i < len(seeds); i++ {
		result := <-results
		locs = append(locs, result)
		fmt.Println("Result", i, ":", result)
	}

	close(results)

	minLoc := slices.Min(locs)

	//fmt.Printf("%+v \n", seeds)
	//fmt.Printf("%+v \n", resources)
	//fmt.Printf("%+v \n", locs)

	fmt.Printf("%+v \n", minLoc)

	return minLoc
}

func getDestinationBySeedV2(seed Seed, resources [][]Resource) int {
	minLoc := math.MaxInt
	ranges := []Range{{seed.Start, seed.Start + seed.RangeLength - 1}}
	for _, resourceGroup := range resources {
		sort.Slice(resourceGroup, func(i, j int) bool {
			return resourceGroup[i].Source < resourceGroup[j].Source
		})

		destinations := getDestinationV2(ranges, resourceGroup)
		//fmt.Printf("Destinations %+v \n\n", destinations)
		ranges = destinations
	}
	for _, rang := range ranges {
		if minLoc > rang.Start {
			minLoc = rang.Start
		}
	}
	return minLoc
}

func getDestinationV2(ranges []Range, resourceGroup []Resource) []Range {
	//fmt.Printf("Ranges %+v \n", ranges)
	//fmt.Printf("ResourceGroup %+v \n", resourceGroup)
	dest := []Range{}
	for _, rang := range ranges {
		dest = append(dest, getDestinationV2ByRange(rang, resourceGroup)...)
	}
	return dest
}

func getDestinationV2ByRange(rang Range, resourceGroup []Resource) []Range {
	resources := resourceGroup
	resourcesRanges := []Range{}
	for _, resource := range resources {
		resourcesRanges = append(resourcesRanges, Range{resource.Source, resource.End})
	}
	sort.Slice(resourcesRanges, func(i, j int) bool {
		return resourcesRanges[i].Start > resourcesRanges[j].Start
	})
	sourceRanges := getOverlapByRange(rang, resourcesRanges)
	//fmt.Printf("Overlap %+v \n", sourceRanges)
	var dest []Range

	for _, sourceRange := range sourceRanges {
		resource := findResource(sourceRange, resources)
		offset := sourceRange.Start - resource.Source
		dest = append(dest, Range{resource.Destination + offset, resource.Destination + offset + sourceRange.rangeLength()})
	}

	return dest
}

func findResource(sourceRange Range, resources []Resource) Resource {
	result := Resource{
		Destination: sourceRange.Start,
		Source:      sourceRange.Start,
		RangeLength: sourceRange.rangeLength(),
	}

	for _, resource := range resources {
		if sourceRange.Start >= resource.Source && sourceRange.Start <= resource.End {
			result = resource
			break
		}
	}

	return result
}

func getOverlapByRange(rang Range, resourceRanges []Range) []Range {
	var points []Point
	var overlap []Range

	points = append(points, Point{rang.Start, SStart}, Point{rang.End, SEnd})
	for _, resourceRange := range resourceRanges {
		if resourceRange.Start > rang.Start && resourceRange.Start <= rang.End {
			points = append(points, Point{resourceRange.Start, RStart})
		}
		if resourceRange.End >= rang.Start && resourceRange.End < rang.End {
			points = append(points, Point{resourceRange.End, REnd})
		}
	}
	sort.Slice(points, func(i, j int) bool {
		var comparison bool
		if points[i].Value == points[j].Value {
			comparison = points[i].Type < points[j].Type
		} else {
			comparison = points[i].Value < points[j].Value
		}
		return comparison
	})
	for i := 0; i < len(points)-1; i++ {
		start := points[i]
		end := points[i+1]
		if start.Type == SStart && end.Type == RStart {
			overlap = append(overlap, Range{start.Value, end.Value - 1})
		} else if start.Type == SStart && (end.Type == REnd || end.Type == SEnd) {
			overlap = append(overlap, Range{start.Value, end.Value})
		} else if start.Type == RStart && (end.Type == REnd || end.Type == SEnd) {
			overlap = append(overlap, Range{start.Value, end.Value})
		} else if start.Type == REnd && end.Type == RStart && (end.Value-1) > (start.Value+1) {
			overlap = append(overlap, Range{start.Value + 1, end.Value - 1})
		} else if start.Type == REnd && end.Type == SEnd {
			overlap = append(overlap, Range{start.Value + 1, end.Value})
		} else {
			continue
		}
	}
	return overlap
}

type Range struct {
	Start int
	End   int
}

func (r *Range) rangeLength() int {
	return r.End - r.Start
}

type Point struct {
	Value int
	Type  int
}

const (
	SStart = iota
	RStart
	REnd
	SEnd
)

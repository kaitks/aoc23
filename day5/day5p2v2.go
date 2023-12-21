package main

import (
	"fmt"
	"golang.org/x/exp/maps"
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

	raw, _ := os.ReadFile(filePath)
	data := string(raw)

	sections := strings.Split(data, "\n\n")
	seedRaws := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	seeds := []Seed{}
	resources := []map[int]Resource{}

	for i := 0; i < len(seedRaws); i = i + 2 {
		seedStart, _ := strconv.Atoi(seedRaws[i])
		seedLength, _ := strconv.Atoi(seedRaws[i+1])
		seeds = append(seeds, Seed{seedStart, seedLength})
	}

	for i := 1; i < len(sections); i++ {
		mapData := strings.Split(strings.Split(sections[i], ":\n")[1], "\n")
		resourceMap := map[int]Resource{}
		for _, mapString := range mapData {
			fields := strings.Split(mapString, " ")
			destination, _ := strconv.Atoi(fields[0])
			source, _ := strconv.Atoi(fields[1])
			rangeLength, _ := strconv.Atoi(fields[2])
			resourceMap[source] = Resource{destination, source, rangeLength}
		}
		resources = append(resources, resourceMap)
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

func getDestinationBySeedV2(seed Seed, resources []map[int]Resource) int {
	minLoc := math.MaxInt
	ranges := []Range{{seed.Start, seed.Start + seed.RangeLength}}
	for _, resourceMap := range resources {
		destinations := getDestinationV2(ranges, resourceMap)
		ranges = destinations
	}
	for _, rang := range ranges {
		if minLoc > rang.Start {
			minLoc = rang.Start
		}
	}
	return minLoc
}

func getDestinationV2(ranges []Range, resourceMap map[int]Resource) []Range {
	dest := []Range{}
	for _, rang := range ranges {
		dest = append(dest, getDestinationV2ByRange(rang, resourceMap)...)
	}
	return dest
}

func getDestinationV2ByRange(rang Range, resourceMap map[int]Resource) []Range {
	resources := maps.Values(resourceMap)
	resourcesRanges := []Range{}
	for _, resource := range resources {
		resourcesRanges = append(resourcesRanges, Range{resource.Source, resource.End()})
	}
	sort.Slice(resourcesRanges, func(i, j int) bool {
		return resourcesRanges[i].Start > resourcesRanges[j].Start
	})
	sourceRanges := getOverlapByRange(rang, resourcesRanges)
	var dest []Range

	for _, sourceRange := range sourceRanges {
		resource, exists := resourceMap[sourceRange.Start]
		if !exists {
			resource = Resource{
				Destination: sourceRange.Start,
				Source:      sourceRange.Start,
				RangeLength: sourceRange.rangeLength(),
			}
		}
		dest = append(dest, Range{resource.Destination, resource.Destination + sourceRange.rangeLength()})
	}

	return dest
}

func getOverlapByRange(rang Range, resourceRanges []Range) []Range {
	var overlap []Range
	startRange := rang
	for i, resourceRange := range resourceRanges {
		if startRange.Start > resourceRange.End {
			if i == len(resourceRanges)-1 {
				overlap = append(overlap, startRange)
			}
			continue
		} else if startRange.Start >= resourceRange.Start && startRange.End <= resourceRange.End {
			overlap = append(overlap, startRange)
			break
		} else if startRange.Start >= resourceRange.Start && startRange.End > resourceRange.End {
			overlap = append(overlap, Range{startRange.Start, resourceRange.End})
			startRange = Range{resourceRange.End + 1, startRange.End}
			if i == len(resourceRanges)-1 {
				overlap = append(overlap, startRange)
			}
		} else if startRange.End < resourceRange.Start {
			overlap = append(overlap, startRange)
			break
		} else if startRange.Start < resourceRange.Start {
			overlap = append(overlap, Range{startRange.Start, resourceRange.Start - 1})
			startRange = Range{resourceRange.Start, startRange.End}
			if i == len(resourceRanges)-1 {
				overlap = append(overlap, startRange)
			}
		} else {
			break
		}
	}
	return overlap
}

//func getOverlap(rang Range, resourceRange Range) []Range {
//	var overlap []Range
//	return overlap
//}

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

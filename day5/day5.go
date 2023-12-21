package main

import (
	"fmt"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func day5(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)

	raw, _ := os.ReadFile(filePath)
	data := string(raw)

	sections := strings.Split(data, "\n\n")
	seeds := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	resources := []map[int]Resource{}

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

	locs := []int{}

	for _, seedString := range seeds {
		seed, _ := strconv.Atoi(seedString)
		source := seed
		for _, resourceMap := range resources {
			destination := getDestination(source, resourceMap)
			source = destination
		}
		locs = append(locs, source)
	}

	//fmt.Printf("%+v \n", seeds)
	//fmt.Printf("%+v \n", resources)
	//fmt.Printf("%+v \n", locs)

	minLoc := slices.Min(locs)
	fmt.Printf("%+v \n", minLoc)

	return minLoc
}

func getDestination(source int, resourceMap map[int]Resource) int {
	nearest := 0
	keys := maps.Keys(resourceMap)
	sort.Ints(keys)
	for _, key := range keys {
		if source >= key {
			nearest = key
		} else {
			break
		}
	}
	resource := resourceMap[nearest]
	result := source
	if resource.Source+resource.RangeLength >= source {
		result = resource.Destination + source - resource.Source
	}
	return result
}

type Resource struct {
	Destination int
	Source      int
	RangeLength int
}

func (r *Resource) offset() int {
	return r.Destination - r.Source
}

func (r *Resource) End() int {
	return r.Source + r.RangeLength
}

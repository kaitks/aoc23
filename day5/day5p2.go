package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func day5p2(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)

	raw, _ := os.ReadFile(filePath)
	data := string(raw)

	sections := strings.Split(data, "\n\n")
	seedRaws := strings.Split(strings.Split(sections[0], ": ")[1], " ")
	resources := []map[int]Resource{}
	seeds := []int{}
	for i := 0; i < len(seedRaws); i = i + 2 {
		seedStart, _ := strconv.Atoi(seedRaws[i])
		seedLength, _ := strconv.Atoi(seedRaws[i+1])
		seedRange := makeRange(seedStart, seedStart+seedLength-1)
		seeds = append(seeds, seedRange...)
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

	locs := []int{}

	for _, seed := range seeds {
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

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

package main

import (
	"fmt"
	"math"
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
			result := getDestinationBySeed(seed, resources)
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

func getDestinationBySeed(seed Seed, resources []map[int]Resource) int {
	minLoc := math.MaxInt
	for i := 0; i < seed.RangeLength; i++ {
		value := seed.Start + i
		source := value

		for _, resourceMap := range resources {
			destination := getDestination(source, resourceMap)
			source = destination
		}
		if minLoc > source {
			minLoc = source
		}
	}
	return minLoc
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

type Seed struct {
	Start       int
	RangeLength int
}

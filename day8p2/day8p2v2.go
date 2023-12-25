package day8p2

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

func day8p2v2(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	raw, _ := os.ReadFile(filePath)
	data := string(raw)

	sections := strings.Split(data, "\n\n")

	instructionsStr := sections[0]
	instructions := strings.Split(instructionsStr, "")
	nodesStr := sections[1]
	var nodes []Node
	nodeMap := map[string]Node{}
	var aNodes []Node

	for _, line := range strings.Split(nodesStr, "\n") {
		node := parseLine(line)
		nodes = append(nodes, node)
		nodeMap[node.Value] = node
		if node.LastChar == "A" {
			aNodes = append(aNodes, node)
		}
	}

	var steps []int

	for _, node := range aNodes {
		currentNode := node
		step := 0
		for currentNode.LastChar != "Z" {
			for _, instruction := range instructions {
				step++
				dest := currentNode.move(instruction, &nodeMap)
				currentNode = dest
				if dest.LastChar == "Z" {
					break
				}
			}
		}
		steps = append(steps, step)
	}

	lcmr := 1

	for _, step := range steps {
		lcmr = int(lcm(int64(lcmr), int64(step)))
	}

	fmt.Printf("Total: %+v\n", lcmr)

	return lcmr
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

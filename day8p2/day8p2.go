package day8p2

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func day8p2(fileName string) int {
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

	steps := 0
	currentNodes := aNodes
	zNodeCount := 0
	aNodeCount := len(aNodes)

	for zNodeCount != aNodeCount {
		for _, instruction := range instructions {
			fmt.Printf("Current Nodes: %+v\n", currentNodes)
			zNodeCount = 0
			var destNodes []Node
			steps++
			for _, node := range currentNodes {
				dest := node.move(instruction, &nodeMap)
				destNodes = append(destNodes, dest)
				if dest.LastChar == "Z" {
					zNodeCount++
				}
			}
			currentNodes = destNodes
			if zNodeCount == aNodeCount {
				break
			}
		}
	}

	fmt.Printf("Total: %+v\n", steps)

	return steps
}

func parseLine(line string) Node {
	parts := strings.Fields(reg.ReplaceAllString(line, ""))
	Value := parts[0]
	Left := parts[1]
	Right := parts[2]
	LastChar := Value[len(Value)-1:]
	return Node{Value, Left, Right, LastChar}
}

type Node struct {
	Value    string
	Left     string
	Right    string
	LastChar string
}

func (node *Node) move(instruction string, nodeMap *map[string]Node) Node {
	destinationValue := ""
	if instruction == "L" {
		destinationValue = node.Left
	} else {
		destinationValue = node.Right
	}
	return (*nodeMap)[destinationValue]
}

var reg = regexp.MustCompile(`[=(),]`)

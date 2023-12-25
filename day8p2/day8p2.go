package day8v2

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
	nodes := []Node{}
	nodeMap := map[string]Node{}
	for _, line := range strings.Split(nodesStr, "\n") {
		node := parseLine(line)
		nodes = append(nodes, node)
		nodeMap[node.Value] = node
	}

	currentNode := nodeMap["AAA"]
	zNode := nodeMap["ZZZ"]
	steps := 0

	for currentNode != zNode {
		for _, instruction := range instructions {
			steps++
			currentNode = currentNode.move(instruction, &nodeMap)
			if currentNode == zNode {
				break
			}
		}
	}

	fmt.Printf("Total: %+v \n", steps)

	return steps
}

func parseLine(line string) Node {
	parts := strings.Fields(reg.ReplaceAllString(line, ""))
	Value := parts[0]
	Left := parts[1]
	Right := parts[2]
	return Node{Value, Left, Right}
}

type Node struct {
	Value string
	Left  string
	Right string
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

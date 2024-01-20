package day19p2

import (
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"slices"
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
	sections := strings.Split(data, "\n\n")
	workflowsStr := sections[0]
	workflowMap := map[string]Workflow{}
	for _, workflowStr := range strings.Split(workflowsStr, "\n") {
		parts := strings.Split(workflowStr, "{")
		label := parts[0]
		workflowStr = parts[1][:len(parts[1])-1]
		rulesStr := strings.Split(workflowStr, ",")
		workflowMap[label] = Workflow{label, rulesStr}
	}
	solutions := findSolutions(&workflowMap)
	total := 0
	possibles := []Solution{}
	for _, solution := range solutions {
		solutionMap := map[string][]Equation{}
		for _, equation := range solution {
			equations, exist := solutionMap[equation.category]
			if !exist {
				equations = []Equation{}
			}
			solutionMap[equation.category] = append(equations, equation)
		}
		possibleSolutionMap := map[string]Range{}
		for category, equations := range solutionMap {
			minVal := math.MinInt
			maxVal := math.MaxInt
			for _, equation := range equations {
				if equation.operation == ">" {
					if minVal < equation.target {
						minVal = equation.target
					}
				} else if equation.operation == "<" {
					if maxVal > equation.target {
						maxVal = equation.target
					}
				}
			}
			possibleSolutionMap[category] = Range{minVal, maxVal}
		}
		possibles = append(possibles, possibleSolutionMap)
	}
	for _, possible := range possibles {
		possible.print()
		total += possible.distinctCombination()
	}

	fmt.Printf("Total: %+v\n", total)
	return total
}

func findSolutions(workflowMap *map[string]Workflow) [][]Equation {
	var solutions [][]Equation
	starter, _ := (*workflowMap)["in"]
	accEquations := []Equation{
		{"x", ">", 0}, {"x", "<", 4001},
		{"m", ">", 0}, {"m", "<", 4001},
		{"a", ">", 0}, {"a", "<", 4001},
		{"s", ">", 0}, {"s", "<", 4001},
	}
	findSolution(workflowMap, &solutions, starter.rulesStr, accEquations)
	return solutions
}

func findSolution(workflowMap *map[string]Workflow, solutions *[][]Equation, rulesStr []string, accEquations []Equation) {
	lenRules := len(rulesStr)
	if lenRules == 0 {
		return
	}
	ruleStr := rulesStr[0]
	matches := re.FindStringSubmatch(ruleStr)
	if len(matches) == 5 {
		category := matches[1]
		operatorStr := matches[2]
		ruleValue, _ := strconv.Atoi(matches[3])
		nextLabel := matches[4]
		equation := Equation{category, operatorStr, ruleValue}
		if nextLabel == "A" {
			*solutions = append(*solutions, append(append([]Equation{}, accEquations...), equation))
			findSolution(workflowMap, solutions, rulesStr[1:], append(accEquations, equation.reverse()))
		} else if nextLabel == "R" {
			findSolution(workflowMap, solutions, rulesStr[1:], append(accEquations, equation.reverse()))
		} else {
			if workflow, exist := (*workflowMap)[nextLabel]; exist {
				findSolution(workflowMap, solutions, workflow.rulesStr, append(append([]Equation{}, accEquations...), equation))
			}
			findSolution(workflowMap, solutions, rulesStr[1:], append(accEquations, equation.reverse()))
		}
	} else {
		nextLabel := ruleStr
		if nextLabel == "A" {
			*solutions = append(*solutions, accEquations)
		} else if nextLabel == "R" {
			return
		} else {
			if workflow, exist := (*workflowMap)[nextLabel]; exist {
				findSolution(workflowMap, solutions, workflow.rulesStr, accEquations)
			}
		}
	}
	return
}

type Workflow struct {
	label    string
	rulesStr []string
}

type Equation struct {
	category  string
	operation string
	target    int
}

type Range struct {
	minVal int
	maxVal int
}

func (equation *Equation) reverse() Equation {
	operation := equation.operation
	target := equation.target
	if operation == ">" {
		operation = "<"
		target = target + 1
	} else if operation == "<" {
		operation = ">"
		target = target - 1
	}
	equation.operation = operation
	equation.target = target
	return Equation{equation.category, operation, target}
}

var re = regexp.MustCompile(`^(\w+)([<>])(\d*):(\w+)$`)

type Solution map[string]Range

func (s *Solution) distinctCombination() int {
	return lo.Reduce(maps.Values(*s), func(agg int, rg Range, _ int) int {
		return agg * max(rg.maxVal-rg.minVal-1, 0)
	}, 1)
}

func (s *Solution) print() {
	fmt.Printf("Solution:\n")
	keys := maps.Keys(*s)
	slices.Sort(keys)
	for _, k := range keys {
		v, _ := (*s)[k]
		fmt.Printf("%d < %s < %d\n", v.minVal, k, v.maxVal)
	}
	fmt.Printf("\n\n")
}

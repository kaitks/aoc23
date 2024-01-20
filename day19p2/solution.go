package day19p2

import (
	"fmt"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"os"
	"path/filepath"
	"regexp"
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
	ratingsStr := sections[1]
	var ratings []Rating
	for _, ratingStr := range strings.Split(ratingsStr, "\n") {
		withoutBracket := ratingStr[1 : len(ratingStr)-1]
		rating := Rating{}
		for _, str := range strings.Split(withoutBracket, ",") {
			equations := strings.Split(str, "=")
			value, _ := strconv.Atoi(equations[1])
			rating[equations[0]] = value
		}
		ratings = append(ratings, rating)
	}
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
	for _, solution := range solutions {
		fmt.Printf("Solution: %+v\n", solution)
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
			minVal := 0
			maxVal := 4000
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
		total += lo.Reduce(maps.Values(possibleSolutionMap), func(agg int, item Range, _ int) int {
			return agg * (item.maxVal - item.minVal - 1)
		}, 1)
	}
	//for _, valid := range valids {
	//	total += lo.Sum(maps.Values(valid))
	//}
	fmt.Printf("Total: %+v\n", total)
	return total
}

func findSolutions(workflowMap *map[string]Workflow) [][]Equation {
	var solutions [][]Equation
	starter, _ := (*workflowMap)["in"]
	accEquations := []Equation{
		{"x", ">", 0}, {"x", "<", 4000},
		{"m", ">", 0}, {"m", "<", 4000},
		{"a", ">", 0}, {"a", "<", 4000},
		{"s", ">", 0}, {"s", "<", 4000},
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
			*solutions = append(*solutions, append(accEquations, equation))
			findSolution(workflowMap, solutions, rulesStr[1:], append(append([]Equation{}, accEquations...), equation.reverse()))
		} else if nextLabel == "R" {
			findSolution(workflowMap, solutions, rulesStr[1:], append(accEquations, equation.reverse()))
		} else {
			if workflow, exist := (*workflowMap)[nextLabel]; exist {
				findSolution(workflowMap, solutions, workflow.rulesStr, append(accEquations, equation))
			}
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

type Rating map[string]int

type Workflow struct {
	label    string
	rulesStr []string
}

type Rule struct {
	nextLabel string
	equations []Equation
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

package day19

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
	re := regexp.MustCompile(`^(\w+)([<>])(\d*):(\w+)$`)
	workflowMap := map[string]Workflow{}
	for _, workflowStr := range strings.Split(workflowsStr, "\n") {
		parts := strings.Split(workflowStr, "{")
		label := parts[0]
		workflowStr = parts[1][:len(parts[1])-1]
		var rules []func(Rating) string
		for _, equationStr := range strings.Split(workflowStr, ",") {
			matches := re.FindStringSubmatch(equationStr)
			if len(matches) == 5 {
				category := matches[1]
				operatorStr := matches[2]
				ruleValue, _ := strconv.Atoi(matches[3])
				nextLabel := matches[4]
				var fn func(int, int) bool
				if operatorStr == ">" {
					fn = more
				} else {
					fn = less
				}
				rule := func(rating Rating) string {
					categoryValue, _ := rating[category]
					if fn(categoryValue, ruleValue) {
						return nextLabel
					} else {
						return ""
					}
				}
				rules = append(rules, rule)
			} else {
				rule := func(_ Rating) string {
					return equationStr
				}
				rules = append(rules, rule)
			}
		}
		workflowMap[label] = Workflow{label, rules}
	}
	valids := lo.Filter(ratings, func(rating Rating, _ int) bool {
		return validate(rating, &workflowMap, "in")
	})
	total := 0
	for _, valid := range valids {
		total += lo.Sum(maps.Values(valid))
	}
	fmt.Printf("Total: %+v\n", total)
	return total
}

func validate(rating Rating, workflowMap *map[string]Workflow, label string) bool {
	workflow, _ := (*workflowMap)[label]
	for _, rule := range workflow.rules {
		nextLabel := rule(rating)
		if nextLabel == "A" {
			return true
		} else if nextLabel == "R" {
			return false
		} else if nextLabel == "" {
			continue
		} else {
			return validate(rating, workflowMap, nextLabel)
		}
	}
	return false
}

type Rating map[string]int

type Workflow struct {
	label string
	rules []func(Rating) string
}

func less(a, b int) bool {
	return a < b
}

func more(a, b int) bool {
	return a > b
}

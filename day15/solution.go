package day15

import (
	"fmt"
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
	instructions := strings.Split(data, ",")
	boxMap := map[int]Box{}
	acc := 0
	for _, instruction := range instructions {
		label, operator, focalLength := instructionParse(instruction)
		boxNumber := hash(label)
		if _, exist := boxMap[boxNumber]; !exist {
			boxMap[boxNumber] = Box{boxNumber, []Len{}}
		}
		box, _ := boxMap[boxNumber]
		foundIndex := slices.IndexFunc(box.LensSl, func(len Len) bool {
			return len.Label == label
		})
		lenn := Len{label, focalLength}
		if operator == "=" {
			if foundIndex == -1 {
				box.LensSl = append(box.LensSl, lenn)
			} else {
				box.LensSl[foundIndex] = lenn
			}
		} else {
			if foundIndex != -1 {
				box.LensSl = slices.Delete(box.LensSl, foundIndex, foundIndex+1)
			}
		}
		boxMap[boxNumber] = box
		fmt.Printf("Box result after %s: %+v\n", instruction, box)
	}

	for _, box := range boxMap {
		cur := 0
		for i, lenn := range box.LensSl {
			cur += (box.Number + 1) * (i + 1) * lenn.FocalLength
		}
		acc += cur
	}

	fmt.Printf("Total: %+v\n", acc)
	return acc
}

func hash(str string) int {
	cur := 0
	for _, char := range str {
		cur += int(char)
		cur = (cur * 17) % 256
	}
	//fmt.Printf("The hash value of %s is %d\n", str, cur)
	return cur
}

func instructionParse(instruction string) (string, string, int) {
	// Regex pattern to match the desired format
	re := regexp.MustCompile(`^(\w+)([=-])(\d*)$`)
	parts := re.FindStringSubmatch(instruction)

	if len(parts) == 4 {
		label := parts[1]
		operator := parts[2]
		focalLength, err := strconv.Atoi(parts[3])
		if err != nil {
			focalLength = 0
		}
		return label, operator, focalLength
	} else {
		fmt.Println("Invalid string format")
		return "", "", 0
	}
}

type Box struct {
	Number int
	//Lens   deque.Deque[Len]
	LensSl []Len
}

type Len struct {
	Label       string
	FocalLength int
}

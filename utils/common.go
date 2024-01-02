package utils

import (
	"fmt"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func ReplaceStringAtIndex(str string, index int, replacement string) string {
	if index >= 0 && index < len(str) {
		return str[:index] + replacement + str[index+1:] // Combine parts
	} else {
		return str // Handle invalid index
	}
}

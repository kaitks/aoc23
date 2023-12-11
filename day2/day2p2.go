package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	var fileName = "input" // Open the file for reading
	file, err := os.Open(filepath.Join(pwd, fileName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	acc := 0

	// Loop over the lines in the file
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		picks := strings.Split(parts[1], "; ")
		red, green, blue := 0, 0, 0
		for _, pickString := range picks {
			cubes := strings.Split(pickString, ", ")
			for _, cubeString := range cubes {
				countString := strings.Split(cubeString, " ")
				count, _ := strconv.Atoi(countString[0])
				switch countString[1] {
				case "red":
					if count > red {
						red = count
					}
				case "green":
					if count > green {
						green = count
					}
				case "blue":
					if count > blue {
						blue = count
					}
				default:
					println(fmt.Errorf("color not defined: %w", countString[1]))
				}
			}
		}
		pick := Pick{red, green, blue}
		acc += pick.Red * pick.Green * pick.Blue
		fmt.Printf("%+v\n", pick)
	}

	println(acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

type Pick struct {
	Red   int
	Green int
	Blue  int
}

func (p Pick) Compare(other Pick) int {
	// Check if all fields are less than
	if p.Red < other.Red && p.Green < other.Green && p.Blue < other.Blue {
		return -1
	}

	// Check if all fields are equal
	if p.Red > other.Red || p.Green > other.Green || p.Blue > other.Blue {
		return 1
	}

	// Otherwise, at least one field must be greater
	return 0
}

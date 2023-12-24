package day7

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func day7(fileName string) int {
	pwd, _ := os.Getwd()
	// Get the file name from the command line argument
	filePath := filepath.Join(pwd, fileName)
	println("Input file:", filePath)
	println("")

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	acc := 0

	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()
		hands = append(hands, parseLine(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		first := hands[i]
		second := hands[j]
		result := true
		for index := 0; index < len(first.KindStrength); index++ {
			if first.KindStrength[index] == second.KindStrength[index] {
				continue
			} else {
				result = first.KindStrength[index] < second.KindStrength[index]
				break
			}
		}
		return result
	})

	for i, hand := range hands {
		rank := i + 1
		fmt.Printf("Rank: %d, Hand: %s, HandStrength: %+v\n", rank, hand.Value, hand.KindStrength)
		acc += rank * hand.Bid
	}

	fmt.Printf("Total: %+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}

func toCard(Value string) Card {
	Strength := ValueStrengthMap[Value]
	return Card{Value, Strength}
}

func toHand(Value string, Cards []Card, Bid int) Hand {
	cardMap := map[Card]int{}
	for _, card := range Cards {
		found, exists := cardMap[card]
		if !exists {
			cardMap[card] = 1
		} else {
			cardMap[card] = found + 1
		}
	}
	inverseCardMap := map[int]Card{}
	for k, v := range cardMap {
		inverseCardMap[v] = k
	}
	counts := maps.Values(cardMap)
	slices.Sort(counts)
	slices.Reverse(counts)

	KeyStrength := []int{}
	maxCount := counts[0]
	if maxCount == 5 {
		KeyStrength = append(KeyStrength, F5AK)
	} else if maxCount == 4 {
		KeyStrength = append(KeyStrength, F4AK)
	} else if maxCount == 3 {
		secondCount := counts[1]
		if secondCount == 2 {
			KeyStrength = append(KeyStrength, FH)
		} else {
			KeyStrength = append(KeyStrength, F3AK)
		}
	} else if maxCount == 2 {
		secondCount := counts[1]
		if secondCount == 2 {
			KeyStrength = append(KeyStrength, TP)
		} else {
			KeyStrength = append(KeyStrength, OP)
		}
	} else {
		KeyStrength = append(KeyStrength, HC)
	}

	for _, card := range Cards {
		KeyStrength = append(KeyStrength, card.Strength)
	}

	return Hand{Value, Cards, Bid, KeyStrength}
}

func parseLine(line string) Hand {
	parts := strings.Fields(line)
	Value := parts[0]
	bidStr := parts[1]
	Cards := []Card{}
	for _, str := range strings.Split(Value, "") {
		Cards = append(Cards, toCard(str))
	}
	Bid, _ := strconv.Atoi(bidStr)
	return toHand(Value, Cards, Bid)
}

type Hand struct {
	Value        string
	Cards        []Card
	Bid          int
	KindStrength []int
}

type Card struct {
	Value    string
	Strength int
}

var ValueStrengthMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

const (
	HC = iota
	OP
	TP
	F3AK
	FH
	F4AK
	F5AK
)

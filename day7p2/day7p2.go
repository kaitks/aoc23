package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func day7p2(fileName string) int {
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
		fmt.Printf("Rank: %d, OriginalHand: %s, Hand: %s, HandStrength: %+v\n", rank, hand.OriginalValue, hand.Value, hand.KindStrength)
		acc += rank * hand.Bid
	}

	fmt.Printf("Total: %+v \n", acc)

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return acc
}

func toHand(Value string, OriginalValue string, Bid int) Hand {
	OriginalCards := toCards(OriginalValue)
	Cards := toCards(Value)
	cardStats, cardMap := toCardStats(Cards)

	stat1 := cardStats[0]
	stat2 := stat1
	if len(cardStats) > 1 {
		stat2 = cardStats[1]
	}
	count1 := stat1.Count
	count2 := stat2.Count
	card1 := stat1.Card
	card2 := stat2.Card

	TypeStrength := 0
	jCount, exists := cardMap[JCard]
	KindStrength := []int{}

	if !exists {
		jCount = 0
	}

	if jCount == 0 || jCount == 5 {
		if count1 == 5 {
			TypeStrength = F5AK
		} else if count1 == 4 {
			TypeStrength = F4AK
		} else if count1 == 3 {
			if count2 == 2 {
				TypeStrength = FH
			} else {
				TypeStrength = F3AK
			}
		} else if count1 == 2 {
			if count2 == 2 {
				TypeStrength = TP
			} else {
				TypeStrength = OP
			}
		} else {
			TypeStrength = HC
		}
		KindStrength = append(KindStrength, TypeStrength)
		for _, card := range OriginalCards {
			KindStrength = append(KindStrength, card.Strength)
		}
		return Hand{Value, OriginalValue, Bid, KindStrength}
	} else {
		if card1 == JCard {
			Value = strings.ReplaceAll(Value, JCard.Value, card2.Value)
			return toHand(Value, OriginalValue, Bid)
		} else {
			Value = strings.ReplaceAll(Value, JCard.Value, card1.Value)
			return toHand(Value, OriginalValue, Bid)
		}
	}
}

func toCard(Value string) Card {
	Strength := ValueStrengthMap[Value]
	return Card{Value, Strength}
}

func toCards(Value string) []Card {
	var Cards []Card
	for _, str := range strings.Split(Value, "") {
		Cards = append(Cards, toCard(str))
	}
	return Cards
}

func toCardStats(Cards []Card) ([]CardStat, map[Card]int) {
	cardMap := map[Card]int{}
	for _, card := range Cards {
		found, exists := cardMap[card]
		if !exists {
			cardMap[card] = 1
		} else {
			cardMap[card] = found + 1
		}
	}
	var cardStats []CardStat
	for card, count := range cardMap {
		cardStats = append(cardStats, CardStat{card, count})
	}
	slices.SortFunc(cardStats, func(a, b CardStat) int {
		if a.Count > b.Count {
			return 1
		} else if a.Count < b.Count {
			return -1
		} else {
			if a.Card.Strength > b.Card.Strength {
				return 1
			} else {
				return -1
			}
		}
	})
	slices.Reverse(cardStats)
	return cardStats, cardMap
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
	return toHand(Value, Value, Bid)
}

type Hand struct {
	Value         string
	OriginalValue string
	Bid           int
	KindStrength  []int
}

type Card struct {
	Value    string
	Strength int
}

type CardStat struct {
	Card  Card
	Count int
}

var JCard = toCard("J")

var ValueStrengthMap = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
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

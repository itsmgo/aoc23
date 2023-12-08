package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var CARDS = map[string]int{
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
	"*": 1,
}

type HandType int

const (
	Five  HandType = 10
	Four  HandType = 8
	Full  HandType = 5
	Three HandType = 3
	Two   HandType = 2
	One   HandType = 1
	High  HandType = 0
)

type Hand struct {
	handType HandType
	hand     string
	bid      int
}

func Solve1(input string) int {
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		hand := strings.Split(line, " ")[0]
		handType := GetHandType(hand)
		bidText := strings.Split(line, " ")[1]
		bid, err := strconv.Atoi(bidText)
		if err != nil {
			fmt.Println("Error converting bid ", bidText)
		}
		hands = append(hands, Hand{handType, hand, bid})
	}

	slices.SortFunc(hands, SortHands)
	solution := 0
	for i, hand := range hands {
		solution += hand.bid * (i + 1)
	}
	return solution
}

func GetHandType(hand string) HandType {
	times := map[string]int{}
	for _, char := range hand {
		_, exists := times[string(char)]
		if exists {
			times[string(char)] += 1
		} else {
			times[string(char)] = 1
		}
	}
	maxCount := 0
	for _, value := range times {
		if value > maxCount {
			maxCount = value
		}
	}
	switch maxCount {
	case 1:
		return High
	case 2:
		if len(times) == 3 {
			return Two
		} else {
			return One
		}
	case 3:
		if len(times) == 2 {
			return Full
		} else {
			return Three
		}
	case 4:
		return Four
	case 5:
		return Five
	}
	fmt.Println("Shouldn't get here counting:", hand)
	return High
}

func SortHands(a Hand, b Hand) int {
	if a.handType != b.handType {
		return int(a.handType - b.handType)
	}

	diff := 0
	for i, char := range a.hand {
		diff = CARDS[string(char)] - CARDS[string(b.hand[i])]
		if diff != 0 {
			return diff
		}
	}
	fmt.Println("Shouldn't get here unless equal:", a, "vs.", b)
	return 0
}

func Solve2(input string) int {
	input = strings.ReplaceAll(input, "J", "*")
	lines := strings.Split(input, "\n")
	hands := make([]Hand, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		hand := strings.Split(line, " ")[0]
		handType := GetHandTypeWithJoker(hand)
		bidText := strings.Split(line, " ")[1]
		bid, err := strconv.Atoi(bidText)
		if err != nil {
			fmt.Println("Error converting bid ", bidText)
		}
		hands = append(hands, Hand{handType, hand, bid})
	}

	slices.SortFunc(hands, SortHands)
	solution := 0
	for i, hand := range hands {
		solution += hand.bid * (i + 1)
	}
	return solution
}

func GetHandTypeWithJoker(hand string) HandType {
	jokers := strings.Count(hand, "*")
	shortHand := strings.ReplaceAll(hand, "*", "")
	times := map[string]int{}
	for _, char := range shortHand {
		_, exists := times[string(char)]
		if exists {
			times[string(char)] += 1
		} else {
			times[string(char)] = 1
		}
	}
	maxCount := 0
	for _, value := range times {
		if value > maxCount {
			maxCount = value
		}
	}
	maxCount += jokers
	switch maxCount {
	case 1:
		return High
	case 2:
		if len(times) == 3 {
			return Two
		} else {
			return One
		}
	case 3:
		if len(times) == 2 {
			return Full
		} else {
			return Three
		}
	case 4:
		return Four
	case 5:
		return Five
	}
	fmt.Println("Shouldn't get here counting:", hand)
	return High
}

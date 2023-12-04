package day4

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Solve1(input string) int {
	cards := strings.Split(input, "\n")
	solution := 0
	for _, card := range cards {
		if card == "" {
			continue
		}
		winning := strings.Split(strings.Split(strings.Split(card, " | ")[0], ": ")[1], " ")
		owned := strings.Split(strings.Split(card, " | ")[1], " ")
		winingNumbers := make([]int, 0)
		ownedNumbers := make([]int, 0)
		for _, win := range winning {
			if win == "" {
				continue
			}
			num, err := strconv.Atoi(win)
			if err != nil {
				fmt.Println("Error converting winning to number '", num, "'")
				return 0
			}
			winingNumbers = append(winingNumbers, num)
		}
		for _, own := range owned {
			if own == "" {
				continue
			}
			num, err := strconv.Atoi(own)
			if err != nil {
				fmt.Println("Error converting owned to number '", num, "-")
				return 0
			}
			ownedNumbers = append(ownedNumbers, num)
		}
		coincidences := 0
		for _, own := range ownedNumbers {
			if IsInArray(own, winingNumbers) {
				coincidences += 1
			}
		}
		if coincidences > 0 {
			solution += int(math.Pow(2, float64(coincidences-1)))
		}
	}
	return solution
}

func IsInArray(target int, numbers []int) bool {
	for _, num := range numbers {
		if num == target {
			return true
		}
	}
	return false
}

func Solve2(input string) int {
	cards := strings.Split(input, "\n")
	instances := make([]int, len(cards)-1)
	for i := range instances {
		instances[i] = 1
	}
	for _, card := range cards {
		if card == "" {
			continue
		}
		cardIds := strings.Split(strings.Split(strings.Split(card, " | ")[0], ": ")[0], " ")
		textId := cardIds[len(cardIds)-1]
		cardId, err := strconv.Atoi(textId)
		if err != nil {
			fmt.Println("Error converting card ID to number '", cardId, "'")
			return 0
		}
		winning := strings.Split(strings.Split(strings.Split(card, " | ")[0], ": ")[1], " ")
		owned := strings.Split(strings.Split(card, " | ")[1], " ")
		winingNumbers := make([]int, 0)
		ownedNumbers := make([]int, 0)
		for _, win := range winning {
			if win == "" {
				continue
			}
			num, err := strconv.Atoi(win)
			if err != nil {
				fmt.Println("Error converting to number '", num, "'")
				return 0
			}
			winingNumbers = append(winingNumbers, num)
		}
		for _, own := range owned {
			if own == "" {
				continue
			}
			num, err := strconv.Atoi(own)
			if err != nil {
				fmt.Println("Error converting to number '", num, "-")
				return 0
			}
			ownedNumbers = append(ownedNumbers, num)
		}
		coincidences := 0
		for _, own := range ownedNumbers {
			if IsInArray(own, winingNumbers) {
				coincidences += 1
			}
		}
		for i := 0; i < coincidences; i++ {
			instances[cardId+i] += instances[cardId-1]
		}
	}
	solution := 0
	for _, num := range instances {
		solution += num
	}
	return solution
}

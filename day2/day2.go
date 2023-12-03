package day2

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func Solve1(input string) int {
	games := strings.Split(input, "\n")
	solution := 0
	for _, game := range games {
		if game == "" {
			continue
		}
		solution += SolveLine(game)
	}
	return solution
}

func SolveLine(line string) int {
	gameId := strings.Split(strings.Split(line, ": ")[0], " ")[1]
	sets := strings.Split(strings.Split(line, ": ")[1], "; ")
	for _, set := range sets {
		groups := strings.Split(set, ", ")
		combination := []int{0, 0, 0}
		for _, group := range groups {
			number_color := strings.Split(group, " ")
			num, err := strconv.Atoi(number_color[0])
			if err != nil {
				fmt.Println("Error converting number ", number_color[0])
			}
			if number_color[1] == "blue" {
				combination[2] = num
			}
			if number_color[1] == "green" {
				combination[1] = num
			}
			if number_color[1] == "red" {
				combination[0] = num
			}
		}
		if combination[0] > MAX_RED || combination[1] > MAX_GREEN || combination[2] > MAX_BLUE {
			return 0
		}
	}
	num, err := strconv.Atoi(gameId)
	if err != nil {
		fmt.Println("Error converting game ID ", gameId)
	}
	return num
}

func Solve2(input string) int {
	games := strings.Split(input, "\n")
	solution := 0
	for _, game := range games {
		if game == "" {
			continue
		}
		solution += SolveLine2(game)
	}
	return solution
}

func SolveLine2(line string) int {
	sets := strings.Split(strings.Split(line, ": ")[1], "; ")
	combination := []int{0, 0, 0}
	for _, set := range sets {
		groups := strings.Split(set, ", ")
		for _, group := range groups {
			number_color := strings.Split(group, " ")
			num, err := strconv.Atoi(number_color[0])
			if err != nil {
				fmt.Println("Error converting number ", number_color[0])
			}
			if number_color[1] == "blue" && num > combination[2] {
				combination[2] = num
			}
			if number_color[1] == "green" && num > combination[1] {
				combination[1] = num
			}
			if number_color[1] == "red" && num > combination[0] {
				combination[0] = num
			}
		}
	}
	return combination[0] * combination[1] * combination[2]
}

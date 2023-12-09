package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve1(input string) int {
	lines := strings.Split(input, "\n")
	solution := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		solution += FindNextNum(line)
	}
	return solution
}

func FindNextNum(line string) int {
	nums := strings.Split(line, " ")
	series := make([]int, 0)
	for _, num := range nums {
		number, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error converting number ", num)
		}
		series = append(series, number)
	}
	lastNums := make([]int, 0)
	lastNums = append(lastNums, series[len(series)-1])
	for {
		diff := make([]int, 0)
		for i, num := range series {
			if i == len(series)-1 {
				break
			}
			diff = append(diff, series[i+1]-num)
		}
		lastNums = append(lastNums, diff[len(diff)-1])
		allZero := true
		for _, val := range diff {
			if val != 0 {
				allZero = false
				break
			}
		}
		if allZero {
			break
		}
		series = diff
	}
	predictions := make([]int, 0)
	predictions = append(predictions, 0)
	j := 0
	for i := len(lastNums) - 2; i >= 0; i-- {
		predictions = append(predictions, lastNums[i]+predictions[j])
		j++
	}
	return predictions[len(predictions)-1]
}

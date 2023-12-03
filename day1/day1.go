package day1

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Solve1(input string) int {
	lines := strings.Split(input, "\n")
	calibration := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		calibration += SolveLine(line)
	}
	return calibration
}

func SolveLine(line string) int {
	calibration := make([]string, 2)
	first := true
	second := false
	for _, char := range line {
		if unicode.IsDigit(char) {
			if first {
				calibration[0] = string(char)
				first = false
			} else {
				calibration[1] = string(char)
				second = true
			}
		}
	}
	if !second {
		calibration[1] = calibration[0]
	}
	num, err := strconv.Atoi(strings.Join(calibration, ""))
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}
	return num
}

func Solve2(input string) int {
	lines := strings.Split(input, "\n")
	calibration := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		line = TransformToDigits(line)
		calibration += SolveLine(line)
	}
	return calibration
}

func TransformToDigits(line string) string {
	text := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digits := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			index := strings.Index(line, text[i])
			if index != -1 {
				line = line[:index+1] + digits[i] + line[index+1:]
			}
		}
	}
	return line
}

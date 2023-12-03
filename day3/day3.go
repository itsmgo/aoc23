package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	x int
	y int
}

type PartNumber struct {
	number int
	ul     Point
	lr     Point
}

func Solve1(input string) int {
	lines := strings.Split(input, "\n")
	symbolsCoords := make([]Point, 0)
	partNumbers := make([]PartNumber, 0)
	for i, line := range lines {
		if line == "" {
			continue
		}
		creatingDigit := false
		ul := Point{0, 0}
		digits := make([]string, 0)
		for j, char := range line {
			if char == '.' {
				creatingDigit, digits, partNumbers = CreateDigit(creatingDigit, digits, partNumbers, ul, i, j)
				continue
			}
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
				if !creatingDigit {
					ul = Point{i - 1, j - 1}
					creatingDigit = true
				}
				continue
			}
			creatingDigit, digits, partNumbers = CreateDigit(creatingDigit, digits, partNumbers, ul, i, j)
			symbolsCoords = append(symbolsCoords, Point{i, j})
		}
		creatingDigit, digits, partNumbers = CreateDigit(creatingDigit, digits, partNumbers, ul, i, 200)
	}
	solution := 0
	for _, pn := range partNumbers {
		isValid := false
		for _, sym := range symbolsCoords {
			if pn.ul.x <= sym.x && pn.ul.y <= sym.y && sym.x <= pn.lr.x && sym.y <= pn.lr.y {
				isValid = true
				break
			}
		}
		if isValid {
			solution += pn.number
		}
	}
	return solution
}

func CreateDigit(creatingDigit bool, digits []string, partNumbers []PartNumber, ul Point, i int, j int) (bool, []string, []PartNumber) {
	if creatingDigit {
		creatingDigit = false
		num, err := strconv.Atoi(strings.Join(digits, ""))
		if err != nil {
			fmt.Println("Error:", err)
		}
		partNumbers = append(partNumbers, PartNumber{num, ul, Point{i + 1, j}})
		digits = make([]string, 0)
	}
	return creatingDigit, digits, partNumbers
}

func Solve2(input string) int {
	lines := strings.Split(input, "\n")
	gearsCoords := make([]Point, 0)
	partNumbers := make([]PartNumber, 0)
	for i, line := range lines {
		if line == "" {
			continue
		}
		creatingDigit := false
		ul := Point{0, 0}
		digits := make([]string, 0)
		for j, char := range line {
			if char == '.' {
				creatingDigit, digits, partNumbers = CreateDigit(creatingDigit, digits, partNumbers, ul, i, j)
				continue
			}
			if unicode.IsDigit(char) {
				digits = append(digits, string(char))
				if !creatingDigit {
					ul = Point{i - 1, j - 1}
					creatingDigit = true
				}
				continue
			}
			if char == '*' {
				creatingDigit, digits, partNumbers = CreateDigit(creatingDigit, digits, partNumbers, ul, i, j)
				gearsCoords = append(gearsCoords, Point{i, j})
			}
		}
		creatingDigit, digits, partNumbers = CreateDigit(creatingDigit, digits, partNumbers, ul, i, 200)
	}
	solution := 0
	for _, gear := range gearsCoords {
		isAdjacent := 0
		gearRatio := 1
		for _, pn := range partNumbers {
			if pn.ul.x <= gear.x && pn.ul.y <= gear.y && gear.x <= pn.lr.x && gear.y <= pn.lr.y {
				isAdjacent += 1
				gearRatio *= pn.number
			}
		}
		if isAdjacent == 2 {
			solution += gearRatio
		}
	}
	return solution
}

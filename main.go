package main

import (
	"fmt"

	"github.com/itsmgo/aoc23/common"
	"github.com/itsmgo/aoc23/day1"
	"github.com/itsmgo/aoc23/day2"
)

func main() {
	input := common.LoadInputContent("day1/input.txt")
	fmt.Println("Day 1, part 1:", day1.Solve1(input))
	fmt.Println("Day 1, part 2:", day1.Solve2(input))

	input = common.LoadInputContent("day2/input.txt")
	fmt.Println("Day 2, part 1:", day2.Solve1(input))
	fmt.Println("Day 2, part 2:", day2.Solve2(input))
}

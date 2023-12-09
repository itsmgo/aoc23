package main

import (
	"fmt"

	"github.com/itsmgo/aoc23/common"
	"github.com/itsmgo/aoc23/day1"
	"github.com/itsmgo/aoc23/day2"
	"github.com/itsmgo/aoc23/day3"
	"github.com/itsmgo/aoc23/day4"
	"github.com/itsmgo/aoc23/day5"
	"github.com/itsmgo/aoc23/day6"
	"github.com/itsmgo/aoc23/day7"
	"github.com/itsmgo/aoc23/day8"
	"github.com/itsmgo/aoc23/day9"
)

func main() {
	input := common.LoadInputContent("day1/input.txt")
	fmt.Println("Day 1, part 1:", day1.Solve1(input))
	fmt.Println("Day 1, part 2:", day1.Solve2(input))

	input = common.LoadInputContent("day2/input.txt")
	fmt.Println("Day 2, part 1:", day2.Solve1(input))
	fmt.Println("Day 2, part 2:", day2.Solve2(input))

	input = common.LoadInputContent("day3/input.txt")
	fmt.Println("Day 3, part 1:", day3.Solve1(input))
	fmt.Println("Day 3, part 2:", day3.Solve2(input))

	input = common.LoadInputContent("day4/input.txt")
	fmt.Println("Day 4, part 1:", day4.Solve1(input))
	fmt.Println("Day 4, part 2:", day4.Solve2(input))

	input = common.LoadInputContent("day5/input.txt")
	fmt.Println("Day 5, part 1:", day5.Solve1(input))
	fmt.Println("Day 5, part 2:", day5.Solve2(input))

	input = common.LoadInputContent("day6/input.txt")
	fmt.Println("Day 6, part 1:", day6.Solve1(input))
	fmt.Println("Day 6, part 2:", day6.Solve2(input))

	input = common.LoadInputContent("day7/input.txt")
	fmt.Println("Day 7, part 1:", day7.Solve1(input))
	fmt.Println("Day 7, part 2:", day7.Solve2(input))

	input = common.LoadInputContent("day8/input.txt")
	fmt.Println("Day 8, part 1:", day8.Solve1(input))
	fmt.Println("Day 8, part 2:", day8.Solve2(input))

	input = common.LoadInputContent("day9/input.txt")
	fmt.Println("Day 9, part 1:", day9.Solve1(input))
}

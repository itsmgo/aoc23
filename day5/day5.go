package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type MapRange struct {
	num_origin      int
	num_destination int
	length          int
}

type Mapping struct {
	origin      string
	destination string
	ranges      []MapRange
}

func Solve1(input string) int {
	lines := strings.Split(input, "\n")
	seedsLine := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seeds := make([]int, 0)
	for _, seedText := range seedsLine {
		num, err := strconv.Atoi(seedText)
		if err != nil {
			fmt.Println("Error converting seed text: ", seedText)
		}
		seeds = append(seeds, num)
	}

	mappingsText := strings.Split(input, "\n\n")[1:]

	mappings := make([]Mapping, 0)
	for _, mapText := range mappingsText {
		mapLines := strings.Split(mapText, "\n")
		mapTitle := mapLines[0]
		mapping := Mapping{}
		mapping.origin = strings.Split(mapTitle, "-to-")[0]
		mapping.destination = strings.Split(strings.Split(mapTitle, "-to-")[1], " ")[0]
		for _, rangeLine := range mapLines[1:] {
			if rangeLine == "" {
				continue
			}
			rangeNums := strings.Split(rangeLine, " ")
			mapRange := MapRange{}
			for i, rangeNum := range rangeNums {
				num, err := strconv.Atoi(rangeNum)
				if err != nil {
					fmt.Println("Error converting range num: ", rangeNum)
				}
				switch i {
				case 0:
					mapRange.num_destination = num
				case 1:
					mapRange.num_origin = num
				case 2:
					mapRange.length = num
				}
			}
			mapping.ranges = append(mapping.ranges, mapRange)
		}
		mappings = append(mappings, mapping)
	}

	locations := make([]int, 0)
	for _, seed := range seeds {
		mapping := FindMapping(mappings, "seed")
		for mapping.destination != "location" {
			fmt.Println(mapping.origin, "->", mapping.destination)
			fmt.Printf("%d -> ", seed)
			MapValue(mapping, &seed)
			fmt.Printf("%d\n", seed)
			mapping = FindMapping(mappings, mapping.destination)
		}
		fmt.Println(mapping.origin, "->", mapping.destination)
		fmt.Printf("%d -> ", seed)
		MapValue(mapping, &seed)
		fmt.Printf("%d\n", seed)
		locations = append(locations, seed)
	}

	fmt.Println(locations)
	solution := locations[0]
	for _, location := range locations {
		if location < solution {
			solution = location
		}
	}
	return solution
}

func FindMapping(mappings []Mapping, origin string) Mapping {
	for _, mapping := range mappings {
		if mapping.origin == origin {
			return mapping
		}
	}
	fmt.Println("Error finding mapping ", origin)
	return mappings[0]
}

func MapValue(mapping Mapping, value *int) {
	for _, mapRange := range mapping.ranges {
		if *value < mapRange.num_origin || *value > mapRange.num_origin+mapRange.length {
			continue
		}
		*value = mapRange.num_destination + *value - mapRange.num_origin
		break
	}
}

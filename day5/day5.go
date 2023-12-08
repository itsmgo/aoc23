package day5

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Segment struct {
	start  int
	length int
}

type MapRange struct {
	destination int
	origin      int
	length      int
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

	mappings := CreateMappings(input)

	locations := make([]int, 0)
	mapping := FindMapping(mappings, "seed")
	for _, seed := range seeds {
		for mapping.destination != "location" {
			MapValue(mapping, &seed)
			mapping = FindMapping(mappings, mapping.destination)
		}
		MapValue(mapping, &seed)
		locations = append(locations, seed)
	}

	solution := locations[0]
	for _, location := range locations {
		if location < solution {
			solution = location
		}
	}
	return solution
}

func CreateMappings(input string) []Mapping {
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
					mapRange.destination = num
				case 1:
					mapRange.origin = num
				case 2:
					mapRange.length = num
				}
			}
			mapping.ranges = append(mapping.ranges, mapRange)
		}
		mappings = append(mappings, mapping)
	}
	return mappings
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
		if *value < mapRange.origin || *value > mapRange.origin+mapRange.length {
			continue
		}
		*value = mapRange.destination + *value - mapRange.origin
		break
	}
}

func Solve2(input string) int {
	lines := strings.Split(input, "\n")
	seedsLine := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	initialSegments := make([]Segment, 0)
	seedStart := 0
	for i, seedText := range seedsLine {
		num, err := strconv.Atoi(seedText)
		if err != nil {
			fmt.Println("Error converting seed text: ", seedText)
		}
		if math.Mod(float64(i), 2) == 0 {
			seedStart = num
		} else {
			initialSegments = append(initialSegments, Segment{seedStart, num})
		}
	}

	mappings := CreateMappings(input)

	mapping := FindMapping(mappings, "seed")
	for mapping.destination != "location" {
		initialSegments = SolveSegmentMapping(initialSegments, mapping)
		mapping = FindMapping(mappings, mapping.destination)
	}
	initialSegments = SolveSegmentMapping(initialSegments, mapping)

	locations := make([]int, 0)
	for _, segment := range initialSegments {
		locations = append(locations, segment.start)
	}

	solution := locations[0]
	for _, location := range locations {
		if location < solution {
			solution = location
		}
	}
	return solution
}

func SolveSegmentMapping(initialSegments []Segment, mapping Mapping) []Segment {
	splitSegments := make([]Segment, 0)
	for _, initialSegment := range initialSegments {
		splitSegments = append(splitSegments, SplitSegment(initialSegment, mapping.ranges)...)
	}
	mappedSegments := make([]Segment, 0)
	for _, segment := range splitSegments {
		mappedSegments = append(mappedSegments, MapSegment(mapping, segment))
	}
	initialSegments = make([]Segment, 0)
	for _, segment := range mappedSegments {
		if !IsContained(segment, initialSegments) {
			initialSegments = append(initialSegments, segment)
		}
	}

	return initialSegments
}

func SplitSegment(segment Segment, mapRanges []MapRange) []Segment {
	segments := make([]Segment, 0)
	for _, mapRange := range mapRanges {
		if mapRange.origin+mapRange.length <= segment.start || mapRange.origin >= segment.start+segment.length {
			continue
		}
		//  S     |---------------|
		//  M |============|
		//
		//		    	||
		//		    	\/
		//
		// S1     |========|
		// S2               |-----|
		if mapRange.origin < segment.start && mapRange.origin+mapRange.length < segment.start+segment.length {
			segments = append(segments, GetSegmentFromLimits(segment.start, mapRange.origin+mapRange.length))
			// segments = append(segments, GetSegmentFromLimits(mapRange.origin+mapRange.length+1, segment.start+segment.length))
		}
		//  S     |----------------|
		//  M        |==========|
		//
		//		    	||
		//		    	\/
		//
		// S1     |-|
		// S2        |==========|
		// S3                    |-|
		if mapRange.origin > segment.start && mapRange.origin+mapRange.length < segment.start+segment.length {
			// segments = append(segments, GetSegmentFromLimits(segment.start, mapRange.origin-1))
			segments = append(segments, GetSegmentFromLimits(mapRange.origin, mapRange.origin+mapRange.length))
			// segments = append(segments, GetSegmentFromLimits(mapRange.origin+mapRange.length+1, segment.start+segment.length))
		}
		//  S     |---------------|
		//  M               |============|
		//
		//		        	||
		//		        	\/
		//
		// S1     |--------|
		// S2               |=====|
		if mapRange.origin > segment.start && mapRange.origin+mapRange.length > segment.start+segment.length {
			// segments = append(segments, GetSegmentFromLimits(segment.start, mapRange.origin-1))
			segments = append(segments, GetSegmentFromLimits(mapRange.origin, segment.start+segment.length))
		}
	}
	slices.SortFunc(segments, SortSegments)
	newSegments := make([]Segment, 0)
	for i, splitSeg := range segments {
		end := 0
		if i == len(segments)-1 {
			end = segment.start + segment.length + 1
		} else {
			end = segments[i+1].start
		}
		if i == 0 && splitSeg.start-segment.start > 1 {
			newSegments = append(newSegments, GetSegmentFromLimits(segment.start, splitSeg.start-1))
		}
		if end-splitSeg.start+splitSeg.length > 1 {
			newSegments = append(newSegments, GetSegmentFromLimits(splitSeg.start+splitSeg.length+1, end-1))
		}
	}
	segments = append(segments, newSegments...)
	if len(segments) == 0 {
		return append(segments, segment)
	}
	return segments
}

func SortSegments(a Segment, b Segment) int {
	return a.start - b.start
}

func MapSegment(mapping Mapping, segment Segment) Segment {
	newSegment := Segment{segment.start, segment.length}
	for _, mapRange := range mapping.ranges {
		if segment.start == mapRange.origin && segment.length > mapRange.length {
			fmt.Println("Error spliting ", segment, "in mapping ", mapRange)
		}
		if segment.start+segment.length == mapRange.origin+mapRange.length && segment.length > mapRange.length {
			fmt.Println("Error spliting ", segment, "in map range ", mapRange, "in mapping ", mapping.origin)
		}
		if segment.start < mapRange.origin || segment.start+segment.length > mapRange.origin+mapRange.length {
			continue
		}
		newSegment.start = mapRange.destination + segment.start - mapRange.origin
	}
	return newSegment
}

func IsContained(segment Segment, totalSegments []Segment) bool {
	for _, tSeg := range totalSegments {
		if segment.start >= tSeg.start && segment.start+segment.length <= tSeg.start+tSeg.length {
			return true
		}
	}
	return false
}

func GetSegmentFromLimits(start int, end int) Segment {
	// if end-start < 0 {
	// 	fmt.Println("Error creating segment: ", start, end)
	// }
	return Segment{start, end - start}
}

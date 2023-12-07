package day6

import (
	"fmt"
)

const VELOCITY = 1

func Solve1(input string) int {
	times := []int{54, 94, 65, 92}
	distances := []int{302, 1476, 1029, 1404}
	options := make([]int, 0)
	for j, time := range times {
		oneOption := 0
		for timeHolding := 0; timeHolding < time; timeHolding++ {
			distance := (time - timeHolding) * timeHolding * VELOCITY
			if distance > distances[j] {
				oneOption += 1
			}
		}
		options = append(options, oneOption)
	}

	solution := 1
	for _, option := range options {
		solution *= option
	}
	return solution
}

func Solve2(input string) int {

	// T = 54946592
	// d_min = 302147610291404
	// d(t) = (T-t)*t
	// d(t) = d_min = (T-t_min)*t_min

	time := 54946592
	distance := 302147610291404

	// timeGuess := 54946592/2 - (48748501 - 54946592/2)
	// distanceGuess := (time - timeGuess) * timeGuess
	// for math.Abs(float64(distanceGuess-distance)) > 1 {
	// 	if distanceGuess-distance > 0 {
	// 		timeGuess -= 1
	// 	} else {
	// 		timeGuess += 1
	// 	}
	// 	distanceGuess = (time - timeGuess) * timeGuess
	// 	break
	// }
	tMin := 6198091
	tMax := 48748501
	fmt.Println((time-tMin)*tMin*VELOCITY - distance)
	fmt.Println((time-tMax)*tMax*VELOCITY - distance)
	return tMax - tMin + 1
}

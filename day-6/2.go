package main

import (
	"fmt"
)

func calculateWays(time, distance int) int {
	ways := 0
	for holdTime := 1; holdTime < time; holdTime++ {
		speed := holdTime
		remainingTime := time - holdTime
		totalDistance := speed * remainingTime
		if totalDistance > distance {
			ways++
		}
	}
	return ways
}

func main() {
	time := 46857582
	distance := 208141212571410

	ways := calculateWays(time, distance)

	fmt.Printf("Ways to win the race: %d\n", ways)
}

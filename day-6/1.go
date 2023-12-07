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
	times := []int{46, 85, 75, 82}
	distances := []int{208, 1412, 1257, 1410}

	totalWays := 1
	for i := range times {
		ways := calculateWays(times[i], distances[i])
		totalWays *= ways
		fmt.Printf("Race %d: %d ways\n", i+1, ways)
	}

	fmt.Printf("Total ways to win all races: %d\n", totalWays)
}

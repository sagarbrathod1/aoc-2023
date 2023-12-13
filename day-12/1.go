package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("puzzle.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalArrangements := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		dots := strings.Split(parts[0], "")
		blockParts := strings.Split(parts[1], ",")
		blocks := make([]int, len(blockParts))
		for i, b := range blockParts {
			fmt.Sscanf(b, "%d", &blocks[i])
		}
		dp := make(map[dpKey]int)
		totalArrangements += countArrangements(dots, blocks, 0, 0, 0, dp)
	}

	fmt.Println("Total arrangements:", totalArrangements)
}

type dpKey struct {
	i, bi, current int
}

func countArrangements(dots []string, blocks []int, i, bi, current int, dp map[dpKey]int) int {
	key := dpKey{i, bi, current}
	if val, exists := dp[key]; exists {
		return val
	}

	if i == len(dots) {
		if bi == len(blocks) && current == 0 {
			return 1
		} else if bi == len(blocks)-1 && blocks[bi] == current {
			return 1
		} else {
			return 0
		}
	}

	ans := 0
	for _, c := range []string{".", "#"} {
		if dots[i] == c || dots[i] == "?" {
			if c == "." && current == 0 {
				ans += countArrangements(dots, blocks, i+1, bi, 0, dp)
			} else if c == "." && current > 0 && bi < len(blocks) && blocks[bi] == current {
				ans += countArrangements(dots, blocks, i+1, bi+1, 0, dp)
			} else if c == "#" {
				ans += countArrangements(dots, blocks, i+1, bi, current+1, dp)
			}
		}
	}

	dp[key] = ans
	return ans
}

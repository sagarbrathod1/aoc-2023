package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MAX = 1000000

var grid [][]int
var rows, cols int

type State struct {
    row, col, dir, len int
}

func minHeatLoss() int {
    dp := make([][][][]int, rows)
    for i := range dp {
        dp[i] = make([][][]int, cols)
        for j := range dp[i] {
            dp[i][j] = make([][]int, 4)
            for k := range dp[i][j] {
                dp[i][j][k] = make([]int, 4)
                for l := range dp[i][j][k] {
                    dp[i][j][k][l] = MAX
                }
            }
        }
    }

    var dirs = []struct{ dr, dc int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    dp[0][0][0][0] = 0

    queue := []State{{0, 0, 0, 0}}
    for len(queue) > 0 {
        state := queue[0]
        queue = queue[1:]

        for newDir := 0; newDir < 4; newDir++ {
            if abs(newDir-state.dir) == 2 {
                continue
            }
            newRow, newCol := state.row+dirs[newDir].dr, state.col+dirs[newDir].dc
            if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
                continue
            }
            newLen := state.len
            if newDir == state.dir {
                newLen++
                if newLen > 3 {
                    continue
                }
            } else {
                newLen = 1
            }
            newHeat := dp[state.row][state.col][state.dir][state.len] + grid[newRow][newCol]
            if newHeat < dp[newRow][newCol][newDir][newLen] {
                dp[newRow][newCol][newDir][newLen] = newHeat
                queue = append(queue, State{newRow, newCol, newDir, newLen})
            }
        }
    }

    minLoss := MAX
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            minLoss = min(minLoss, dp[rows-1][cols-1][i][j])
        }
    }
    return minLoss
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func readGridFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
			return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid = [][]int{}
	for scanner.Scan() {
			line := scanner.Text()
			row := make([]int, len(line))
			for i, ch := range strings.TrimSpace(line) {
					row[i] = int(ch - '0')
			}
			grid = append(grid, row)
	}

	if err := scanner.Err(); err != nil {
			return err
	}
	rows = len(grid)
	if rows > 0 {
			cols = len(grid[0])
	}
	return nil
}

func main() {
	filename := "puzzle.txt"

	if err := readGridFromFile(filename); err != nil {
			fmt.Println("Error reading file:", err)
			return
	}

	fmt.Println("Minimum heat loss:", minHeatLoss())
}

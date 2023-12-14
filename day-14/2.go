package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    file, err := os.Open("puzzle.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)

    var grid [][]rune
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        grid = append(grid, []rune(line))
    }

    if len(grid) == 0 {
        fmt.Println("Error: Grid is empty.")
        return
    }

    fmt.Println(calculateLoadAfterCycles(grid, 1000000000))
}

func calculateLoadAfterCycles(grid [][]rune, cycles int) int {
    seen := make(map[string]int)
    for cycle := 0; cycle < cycles; cycle++ {
        for i := 0; i < 4; i++ {
            tilt(grid)
            grid = rotate(grid)
        }

        gridHash := gridToString(grid)

        if lastCycle, found := seen[gridHash]; found {
            cycleLength := cycle - lastCycle
            remainingCycles := (cycles - cycle) / cycleLength
            cycle += remainingCycles * cycleLength
        } else {
            seen[gridHash] = cycle
        }
    }
    return calculateLoad(grid)
}

func tilt(grid [][]rune) {
    for c := range grid[0] {
        for {
            moved := false
            for r := 1; r < len(grid); r++ {
                if grid[r][c] == 'O' && grid[r-1][c] == '.' {
                    grid[r][c], grid[r-1][c] = grid[r-1][c], grid[r][c]
                    moved = true
                }
            }
            if !moved {
                break
            }
        }
    }
}

func rotate(grid [][]rune) [][]rune {
    newGrid := make([][]rune, len(grid[0]))
    for i := range newGrid {
        newGrid[i] = make([]rune, len(grid))
    }
    for r := range grid {
        for c := range grid[r] {
            newGrid[c][len(grid)-r-1] = grid[r][c]
        }
    }
    return newGrid
}

func calculateLoad(grid [][]rune) int {
    load := 0
    for r, row := range grid {
        for _, cell := range row {
            if cell == 'O' {
                load += len(grid) - r
            }
        }
    }
    return load
}

func gridToString(grid [][]rune) string {
    var sb strings.Builder
    for _, row := range grid {
        sb.WriteString(string(row))
        sb.WriteRune('\n')
    }
    return sb.String()
}

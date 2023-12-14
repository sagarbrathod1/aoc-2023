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

    tiltNorth(grid)

    fmt.Println(calculateLoad(grid))
}

func tiltNorth(grid [][]rune) {
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

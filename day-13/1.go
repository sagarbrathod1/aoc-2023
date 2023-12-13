package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    file, err := os.Open("puzzle.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var grids [][]string
    var grid []string

    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            grids = append(grids, grid)
            grid = nil
        } else {
            grid = append(grid, line)
        }
    }
    grids = append(grids, grid)

    answer := 0
    for _, g := range grids {
        R, C := len(g), len(g[0])
        for c := 0; c < C-1; c++ {
            mismatchCount := 0
            for dc := 0; dc < C; dc++ {
                left, right := c-dc, c+1+dc
                if left >= 0 && right < C {
                    for r := 0; r < R; r++ {
                        if g[r][left] != g[r][right] {
                            mismatchCount++
                        }
                    }
                }
            }
            if mismatchCount == 0 {
                answer += c + 1
            }
        }
        for r := 0; r < R-1; r++ {
            mismatchCount := 0
            for dr := 0; dr < R; dr++ {
                up, down := r-dr, r+1+dr
                if up >= 0 && down < R {
                    for c := 0; c < C; c++ {
                        if g[up][c] != g[down][c] {
                            mismatchCount++
                        }
                    }
                }
            }
            if mismatchCount == 0 {
                answer += 100 * (r + 1)
            }
        }
    }
    fmt.Println(answer)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func expandUniverse(grid []string) []string {
    rowCount := len(grid)
    colCount := len(grid[0])

    expandRow := make([]bool, rowCount)
    expandCol := make([]bool, colCount)

    for i := 0; i < rowCount; i++ {
        expandRow[i] = !strings.Contains(grid[i], "#")
    }

    for j := 0; j < colCount; j++ {
        hasGalaxy := false
        for i := 0; i < rowCount; i++ {
            if grid[i][j] == '#' {
                hasGalaxy = true
                break
            }
        }
        expandCol[j] = !hasGalaxy
    }

    newGrid := make([]string, 0)
    for i := 0; i < rowCount; i++ {
        newRow := ""
        for j := 0; j < colCount; j++ {
            newRow += string(grid[i][j])
            if expandCol[j] {
                newRow += string(grid[i][j])
            }
        }
        newGrid = append(newGrid, newRow)
        if expandRow[i] {
            newGrid = append(newGrid, newRow)
        }
    }

    return newGrid
}

func bfs(grid []string, start, end [2]int) int {
    rowCount := len(grid)
    colCount := len(grid[0])
    visited := make([][]bool, rowCount)
    for i := range visited {
        visited[i] = make([]bool, colCount)
    }

    type point struct {
        x, y, dist int
    }

    queue := []point{{start[0], start[1], 0}}
    visited[start[0]][start[1]] = true

    directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    for len(queue) > 0 {
        p := queue[0]
        queue = queue[1:]

        if p.x == end[0] && p.y == end[1] {
            return p.dist
        }

        for _, d := range directions {
            newX, newY := p.x+d[0], p.y+d[1]
            if newX >= 0 && newX < rowCount && newY >= 0 && newY < colCount && !visited[newX][newY] {
                visited[newX][newY] = true
                queue = append(queue, point{newX, newY, p.dist + 1})
            }
        }
    }

		return -1
}

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
			panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan() {
			input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
			panic(err)
	}

	expandedGrid := expandUniverse(input)

	galaxies := make([][2]int, 0)
	for i, row := range expandedGrid {
			for j, cell := range row {
					if cell == '#' {
							galaxies = append(galaxies, [2]int{i, j})
					}
			}
	}

	totalPathLength := 0
	for i := 0; i < len(galaxies)-1; i++ {
			for j := i + 1; j < len(galaxies); j++ {
					pathLength := bfs(expandedGrid, galaxies[i], galaxies[j])
					totalPathLength += pathLength
			}
	}

	fmt.Println("Total Length of Shortest Paths:", totalPathLength)
}

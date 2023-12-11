package main

import (
	"bufio"
	"fmt"
	"os"
)

func neighbors(pos [2]int, grid [][]rune) [][2]int {
    i, j := pos[0], pos[1]
    c := grid[i][j]
    var neighborPositions [][2]int

    switch c {
    case '|':
        neighborPositions = [][2]int{{i + 1, j}, {i - 1, j}}
    case '-':
        neighborPositions = [][2]int{{i, j - 1}, {i, j + 1}}
    case 'L':
        neighborPositions = [][2]int{{i - 1, j}, {i, j + 1}}
    case 'J':
        neighborPositions = [][2]int{{i - 1, j}, {i, j - 1}}
    case '7', 'S':
        neighborPositions = [][2]int{{i + 1, j}, {i, j - 1}}
    case 'F':
        neighborPositions = [][2]int{{i + 1, j}, {i, j + 1}}
    }

    var validNeighbors [][2]int
    for _, v := range neighborPositions {
        if v[0] >= 0 && v[1] >= 0 && v[0] < len(grid) && v[1] < len(grid[0]) {
            validNeighbors = append(validNeighbors, v)
        }
    }

    return validNeighbors
}

func main() {
    file, err := os.Open("puzzle.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var grid [][]rune
    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, []rune(line))
    }

    n, m := len(grid), len(grid[0])
    var start [2]int
    loop := make(map[[2]int]bool)

    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 'S' {
                start = [2]int{i, j}
            }
        }
    }

    queue := [][2]int{start}
    loop[start] = true
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        for _, neighbor := range neighbors(current, grid) {
            if !loop[neighbor] {
                queue = append(queue, neighbor)
                loop[neighbor] = true
            }
        }
    }

    inside := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if !loop[[2]int{i, j}] {
                counter := 0
                for l := 0; l < j; l++ {
                    if loop[[2]int{i, l}] && grid[i][l] != '-' && grid[i][l] != 'J' && grid[i][l] != 'L' {
                        counter++
                    }
                }
                if counter%2 != 0 {
                    inside++
                }
            }
        }
    }

    fmt.Println("Number of tiles inside the loop:", inside)
}

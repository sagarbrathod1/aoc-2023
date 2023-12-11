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

    var start [2]int
    for i, line := range grid {
        for j, r := range line {
            if r == 'S' {
                start = [2]int{i, j}
                break
            }
        }
    }

    queue := [][2]int{start}
    distance := make(map[[2]int]int)
    distance[start] = 0
    maxDistance := 0

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        for _, neighbor := range neighbors(current, grid) {
            if _, ok := distance[neighbor]; !ok {
                queue = append(queue, neighbor)
                distance[neighbor] = distance[current] + 1
                if distance[neighbor] > maxDistance {
                    maxDistance = distance[neighbor]
                }
            }
        }
    }

    fmt.Println("Farthest point from the start is", maxDistance, "steps away.")
}

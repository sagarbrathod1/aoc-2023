package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type State struct {
    dist, row, col, dir, indir int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].dist < pq[j].dist
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func readGridFromFile(filename string) ([]string, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var grid []string
    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, strings.TrimSpace(line))
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }
    return grid, nil
}

func solve(grid []string, part2 bool) int {
    R := len(grid)
    C := len(grid[0])
    dirs := []struct{ dr, dc int }{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, State{0, 0, 0, -1, -1})
    visited := make(map[[4]int]bool)

    for pq.Len() > 0 {
        state := heap.Pop(pq).(State)
        if state.row == R-1 && state.col == C-1 {
            return state.dist
        }

        if visited[[4]int{state.row, state.col, state.dir, state.indir}] {
            continue
        }
        visited[[4]int{state.row, state.col, state.dir, state.indir}] = true

        for i, dir := range dirs {
            newRow, newCol := state.row+dir.dr, state.col+dir.dc
            newDir := i
            newIndir := 1
            if newDir == state.dir {
                newIndir = state.indir + 1
            }

            isntReverse := (newDir+2)%4 != state.dir
            isValidPart1 := newIndir <= 3
            isValidPart2 := newIndir <= 10 && (newDir == state.dir || state.indir >= 4 || state.indir == -1)
            isValid := isValidPart2
            if !part2 {
                isValid = isValidPart1
            }

            if newRow >= 0 && newRow < R && newCol >= 0 && newCol < C && isntReverse && isValid {
                cost := int(grid[newRow][newCol] - '0')
                newState := State{state.dist + cost, newRow, newCol, newDir, newIndir}
                heap.Push(pq, newState)
            }
        }
    }
    return -1
}

func main() {
    filename := "puzzle.txt"
    grid, err := readGridFromFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    fmt.Println("Minimum heat loss (Part 1):", solve(grid, false))
    fmt.Println("Minimum heat loss (Part 2):", solve(grid, true))
}

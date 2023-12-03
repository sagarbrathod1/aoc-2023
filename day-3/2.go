package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func considerNumberNeighbors(board [][]rune, startY, startX, endY, endX int, num int, gearNums map[[2]int][]int) bool {
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			if y >= 0 && y < len(board) && x >= 0 && x < len(board[y]) {
				char := board[y][x]
				if char != '.' && (char < '0' || char > '9') {
					if char == '*' {
						gearNums[[2]int{y, x}] = append(gearNums[[2]int{y, x}], num)
					}
					return true
				}
			}
		}
	}
	return false
}

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var board [][]rune
	for scanner.Scan() {
		board = append(board, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	total := 0
	gearNums := make(map[[2]int][]int)
	numPattern := regexp.MustCompile(`\d+`)

	for y, row := range board {
		matches := numPattern.FindAllStringIndex(string(row), -1)
		for _, match := range matches {
			startX, endX := match[0], match[1]
			num, _ := strconv.Atoi(string(row[startX:endX]))
			if considerNumberNeighbors(board, y-1, startX-1, y+1, endX, num, gearNums) {
				total += num
			}
		}
	}

	fmt.Println("Total:", total)

	ratTotal := 0
	for _, v := range gearNums {
		if len(v) == 2 {
			ratTotal += v[0] * v[1]
		}
	}
	fmt.Println("Rat Total:", ratTotal)
}

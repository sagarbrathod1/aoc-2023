package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	UP    = 0
	RIGHT = 1
	DOWN  = 2
	LEFT  = 3
)

var (
	DR = []int{-1, 0, 1, 0}
	DC = []int{0, 1, 0, -1}
)

func step(r, c, d int) (int, int, int) {
	return r + DR[d], c + DC[d], d
}

func score(grid [][]byte, sr, sc, sd int) int {
	R := len(grid)
	C := len(grid[0])
	pos := []int{sr, sc, sd}
	seen := make(map[int]bool)
	seen2 := make(map[int]bool)

	for len(pos) > 0 {
		var np []int
		for i := 0; i < len(pos); i += 3 {
			r, c, d := pos[i], pos[i+1], pos[i+2]
			if r >= 0 && r < R && c >= 0 && c < C {
				seen[r*C+c] = true
				if seen2[r*C*4+c*4+d] {
					continue
				}
				seen2[r*C*4+c*4+d] = true
				ch := grid[r][c]
				switch ch {
				case '.':
					nr, nc, nd := step(r, c, d)
					np = append(np, nr, nc, nd)
				case '/':
					var nd int
					switch d {
					case UP:
						nd = RIGHT
					case RIGHT:
						nd = UP
					case DOWN:
						nd = LEFT
					case LEFT:
						nd = DOWN
					}
					nr, nc, _ := step(r, c, nd)
					np = append(np, nr, nc, nd)
				case '\\':
					var nd int
					switch d {
					case UP:
						nd = LEFT
					case RIGHT:
						nd = DOWN
					case DOWN:
						nd = RIGHT
					case LEFT:
						nd = UP
					}
					nr, nc, _ := step(r, c, nd)
					np = append(np, nr, nc, nd)
				case '|':
					if d == UP || d == DOWN {
						nr, nc, nd := step(r, c, d)
						np = append(np, nr, nc, nd)
					} else {
						nr, nc, _ := step(r, c, UP)
						np = append(np, nr, nc, UP)
						nr, nc, _ = step(r, c, DOWN)
						np = append(np, nr, nc, DOWN)
					}
				case '-':
					if d == LEFT || d == RIGHT {
						nr, nc, nd := step(r, c, d)
						np = append(np, nr, nc, nd)
					} else {
						nr, nc, _ := step(r, c, LEFT)
						np = append(np, nr, nc, LEFT)
						nr, nc, _ = step(r, c, RIGHT)
						np = append(np, nr, nc, RIGHT)
					}
				default:
					panic("invalid character")
				}
			}
		}
		pos = np
	}
	return len(seen)
}

func main() {
	file, err := os.Open("puzzle.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}

	R := len(grid)
	C := len(grid[0])
	var ans int
	for r := 0; r < R; r++ {
		ans = max(ans, score(grid, r, 0, RIGHT))
		ans = max(ans, score(grid, r, C-1, LEFT))
	}
	for c := 0; c < C; c++ {
		ans = max(ans, score(grid, 0, c, DOWN))
		ans = max(ans, score(grid, R-1, c, UP))
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
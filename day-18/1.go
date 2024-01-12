package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Instruction struct {
	d string
	n int
}

func parseInput(input string) []Instruction {
	var instructions []Instruction
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		var d string
		var n int
		line := scanner.Text()
		if len(line) > 0 {
			fmt.Sscanf(line, "%s %d", &d, &n)
			instructions = append(instructions, Instruction{d: d, n: n})
		}
	}
	return instructions
}

func areaShoelace(cmds []Instruction) int {
	var V []struct{ x, y int }
	pos := struct{ x, y int }{0, 0}
	V = append(V, pos)

	for _, cmd := range cmds {
		switch cmd.d {
		case "R":
			pos.x += cmd.n
		case "D":
			pos.y -= cmd.n
		case "L":
			pos.x -= cmd.n
		case "U":
			pos.y += cmd.n
		}
		V = append(V, pos)
	}

	area := 0
	for i := range V {
		area -= V[i].x * V[(i+1)%len(V)].y
		area += V[i].y * V[(i+1)%len(V)].x
	}
	area /= 2
	return area
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <file>")
		os.Exit(1)
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	cmds := parseInput(string(input))
	perimeter := 0
	for _, cmd := range cmds {
		perimeter += cmd.n
	}
	area := areaShoelace(cmds)
	ans := area + (perimeter/2) + 1
	fmt.Println(ans)
}

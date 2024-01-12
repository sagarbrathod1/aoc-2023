package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	d string
	n int
}

func parseInputPart2(input string) []Instruction {
	var instructions []Instruction
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			hex := line[strings.Index(line, "#")+1 : strings.Index(line, ")")]
			dirCode, _ := strconv.ParseInt(string(hex[5]), 16, 64)
			n, _ := strconv.ParseInt(hex[:5], 16, 64)

			var d string
			switch dirCode {
			case 0:
				d = "R"
			case 1:
				d = "D"
			case 2:
				d = "L"
			case 3:
				d = "U"
			}
			instructions = append(instructions, Instruction{d: d, n: int(n)})
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
	cmds := parseInputPart2(string(input))
	perimeter := 0
	for _, cmd := range cmds {
		perimeter += cmd.n
	}
	area := areaShoelace(cmds)
	ans := area + (perimeter/2) + 1
	fmt.Println(ans)
}

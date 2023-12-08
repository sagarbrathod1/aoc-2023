package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions := scanner.Text()

	network := make(map[string][2]string)
	var startNodes []string
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}
		node := parts[0]
		if strings.HasSuffix(node, "A") {
			startNodes = append(startNodes, node)
		}
		nodes := strings.Split(strings.Trim(parts[1], "()"), ", ")
		if len(nodes) != 2 {
			fmt.Println("Invalid node format for", node, ":", parts[1])
			continue
		}
		network[node] = [2]string{nodes[0], nodes[1]}
	}

	lcm := findLCMSteps(instructions, network, startNodes)
	fmt.Println("LCM of steps to reach all ZZZ:", lcm)
}

func findSteps(node string, instructions string, network map[string][2]string) int {
	steps := 0
	for {
		direction := instructions[steps%len(instructions)]
		nextNode := network[node][0]
		if direction == 'R' {
			nextNode = network[node][1]
		}
		node = nextNode
		steps++
		if strings.HasSuffix(node, "Z") {
			break
		}
	}
	return steps
}

func findLCMSteps(instructions string, network map[string][2]string, startNodes []string) *big.Int {
	lcm := big.NewInt(1)
	for _, node := range startNodes {
		steps := big.NewInt(int64(findSteps(node, instructions, network)))
		gcd := big.NewInt(0)
		gcd.GCD(nil, nil, lcm, steps)
		lcm.Mul(lcm, steps.Div(steps, gcd))
	}
	return lcm
}

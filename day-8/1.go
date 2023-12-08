package main

import (
	"bufio"
	"fmt"
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
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, " = ")
        if len(parts) != 2 {
            fmt.Println("Invalid line format:", line)
            continue
        }
        node := parts[0]
        nodes := strings.Split(strings.Trim(parts[1], "()"), ", ")
        if len(nodes) != 2 {
            fmt.Println("Invalid node format for", node, ":", parts[1])
            continue
        }
        network[node] = [2]string{nodes[0], nodes[1]}
    }

    steps := findZZZ(instructions, network)
    fmt.Println("Steps to reach ZZZ:", steps)
}

func findZZZ(instructions string, network map[string][2]string) int {
    steps := 0
    currentNode := "AAA"
    for i := 0; currentNode != "ZZZ"; i = (i + 1) % len(instructions) {
        if instructions[i] == 'L' {
            currentNode = network[currentNode][0]
        } else {
            currentNode = network[currentNode][1]
        }
        steps++
    }
    return steps
}

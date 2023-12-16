package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
}

func hash(s string) int {
	currentValue := 0
	for _, char := range s {
		asciiValue := int(char)
		currentValue += asciiValue
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

func processStep(boxes [][]Lens, step string) {
	parts := strings.Split(step, "=")
	if len(parts) == 1 {
		parts = strings.Split(step, "-")
	}

	label := parts[0]
	boxIndex := hash(label)

	if strings.Contains(step, "=") {
		focalLength, _ := strconv.Atoi(parts[1])
		found := false
		for i, lens := range boxes[boxIndex] {
			if lens.Label == label {
				boxes[boxIndex][i].FocalLength = focalLength
				found = true
				break
			}
		}
		if !found {
			boxes[boxIndex] = append(boxes[boxIndex], Lens{Label: label, FocalLength: focalLength})
		}
	} else {
		for i, lens := range boxes[boxIndex] {
			if lens.Label == label {
				boxes[boxIndex] = append(boxes[boxIndex][:i], boxes[boxIndex][i+1:]...)
				break
			}
		}
	}
}

func calculateFocusingPower(boxes [][]Lens) int {
	totalPower := 0
	for boxIndex, lenses := range boxes {
		for slotIndex, lens := range lenses {
			power := (boxIndex + 1) * (slotIndex + 1) * lens.FocalLength
			totalPower += power
		}
	}
	return totalPower
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run script.go <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	sequence := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	steps := strings.Split(sequence, ",")
	boxes := make([][]Lens, 256)
	for _, step := range steps {
		processStep(boxes, step)
	}

	fmt.Println(calculateFocusingPower(boxes))
}

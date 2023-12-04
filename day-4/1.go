package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    file, err := os.Open("puzzle.txt")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    totalPoints := 0
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "|")
        if len(parts) != 2 {
            fmt.Println("Invalid line format:", line)
            continue
        }

        winningNumbers := convertToIntSlice(parts[0])
        yourNumbers := convertToIntSlice(parts[1])
        points := calculatePoints(winningNumbers, yourNumbers)
        totalPoints += points
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    } else {
        fmt.Println("Total points:", totalPoints)
    }
}

func convertToIntSlice(s string) []int {
	var nums []int
	for _, numStr := range strings.Fields(s) {
			num, err := strconv.Atoi(numStr)
			if err != nil { // Skip non-numeric strings
					continue
			}
			nums = append(nums, num)
	}
	return nums
}

func calculatePoints(winningNumbers, yourNumbers []int) int {
    points := 0
    for _, yourNum := range yourNumbers {
        for _, winningNum := range winningNumbers {
            if yourNum == winningNum {
                if points == 0 {
                    points = 1
                } else {
                    points *= 2
                }
                break
            }
        }
    }
    return points
}

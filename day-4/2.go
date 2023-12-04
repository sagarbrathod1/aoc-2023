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

    var cards [][][]int // Store winning and your numbers for each card
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, "|")
        if len(parts) != 2 {
            fmt.Println("Invalid line format:", line)
            continue
        }
        cards = append(cards, [][]int{convertToIntSlice(parts[0]), convertToIntSlice(parts[1])})
    }

    totalCards := processCards(cards)

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    } else {
        fmt.Println("Total scratchcards:", totalCards)
    }
}

func convertToIntSlice(s string) []int {
    var nums []int
    for _, numStr := range strings.Fields(s) {
        num, err := strconv.Atoi(numStr)
        if err != nil {
            continue // Skip non-numeric strings
        }
        nums = append(nums, num)
    }
    return nums
}

func processCards(cards [][][]int) int {
    total := len(cards) // Count original cards
    copies := make([]int, len(cards))

    for i, card := range cards {
        matches := countMatches(card[0], card[1])
        for j := 1; j <= matches && i+j < len(cards); j++ {
            copies[i+j]++
        }
    }

    for i := 0; i < len(copies); i++ {
        if copies[i] > 0 {
            total += copies[i] // Add the number of copies for each card
            for j := 1; j <= countMatches(cards[i][0], cards[i][1]) && i+j < len(cards); j++ {
                copies[i+j] += copies[i]
            }
        }
    }

    return total
}

func countMatches(winningNumbers, yourNumbers []int) int {
    matches := 0
    for _, num := range yourNumbers {
        if contains(winningNumbers, num) {
            matches++
        }
    }
    return matches
}

func contains(slice []int, val int) bool {
    for _, item := range slice {
        if item == val {
            return true
        }
    }
    return false
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
    result := solve("input.txt")
    fmt.Println("Result:", result)
}

func solve(path string) int {
    file, err := os.Open(path)
    if err != nil {
        log.Panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    hands := []Hand{}
    for scanner.Scan() {
        line := scanner.Text()
        hand := parseInput(line)
        hands = append(hands, hand)
    }

    sort.Slice(hands, func(i, j int) bool {
        if hands[i].Combo != hands[j].Combo {
            return hands[i].Combo < hands[j].Combo
        }
        for k := 0; k < len(hands[i].Card); k++ {
            if hands[i].Card[k] != hands[j].Card[k] {
                return cardPower(rune(hands[i].Card[k])) < cardPower(rune(hands[j].Card[k]))
            }
        }
        return hands[i].Jokers < hands[j].Jokers
    })

    total := 0
    for rank, hand := range hands {
        total += (rank + 1) * hand.Bid
    }

    return total
}

type Hand struct {
    Card   string
    Bid    int
    Combo  Combo
    Jokers int
}

type Combo int

const (
    FIVE_OF_A_KIND  Combo = 7
    FOUR_OF_A_KIND  Combo = 6
    FULL_HOUSE      Combo = 5
    THREE_OF_A_KIND Combo = 4
    TWO_PAIR        Combo = 3
    ONE_PAIR        Combo = 2
    HIGH_CARD       Combo = 1
)

func parseInput(line string) Hand {
    parts := strings.Split(line, " ")
    card := parts[0]
    bid, err := strconv.Atoi(parts[1])
    if err != nil {
        log.Panic(err)
    }

    counts := make(map[rune]int)
    jokers := 0
    for _, c := range card {
        if c == 'J' {
            jokers++
        } else {
            counts[c]++
        }
    }

    combo := determineCombo(counts, jokers)

    return Hand{
        Card:   card,
        Bid:    bid,
        Combo:  combo,
        Jokers: jokers,
    }
}

func determineCombo(counts map[rune]int, jokers int) Combo {
	// Highest count of the same card
	maxCount := 0
	for _, count := range counts {
			if count > maxCount {
					maxCount = count
			}
	}

	switch maxCount {
	case 4:
			return FIVE_OF_A_KIND
	case 3:
			if jokers > 0 {
					return FIVE_OF_A_KIND
			}
			if len(counts) == 2 || (len(counts) == 3 && jokers > 0) {
					return FULL_HOUSE
			}
			return THREE_OF_A_KIND
	case 2:
			if jokers > 0 {
					if len(counts) == 2 || (len(counts) == 3 && jokers >= 1) {
							return FULL_HOUSE
					}
					return FOUR_OF_A_KIND
			}
			if len(counts) == 3 {
					return TWO_PAIR
			}
			return ONE_PAIR
	case 1:
			if jokers == 4 {
					return FIVE_OF_A_KIND
			}
			if jokers > 0 {
					return THREE_OF_A_KIND
			}
			return HIGH_CARD
	default:
			return HIGH_CARD
	}
}


var cardRank = map[rune]int{
    'A': 14, 'K': 13, 'Q': 12, 'T': 10,
    '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, 'J': 1,
}

func cardPower(c rune) int {
    return cardRank[c]
}

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
		return false
	})	

	total := 0
	for rank, hand := range hands {
		total += (rank + 1) * hand.Bid
	}

	return total
}

type Hand struct {
	Card  string
	Bid   int
	Combo Combo
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
	for _, c := range card {
		counts[c]++
	}

	var combo Combo
	for _, count := range counts {
		switch count {
		case 5:
			combo = FIVE_OF_A_KIND
		case 4:
			combo = FOUR_OF_A_KIND
		case 3:
			if combo != ONE_PAIR { // Three of a kind or full house
				combo = THREE_OF_A_KIND
			} else {
				combo = FULL_HOUSE
			}
		case 2:
			if combo != THREE_OF_A_KIND { // One pair or two pair
				if combo == ONE_PAIR {
					combo = TWO_PAIR
				} else {
					combo = ONE_PAIR
				}
			} else {
				combo = FULL_HOUSE
			}
		default:
			if combo == 0 {
				combo = HIGH_CARD
			}
		}
	}

	return Hand{
		Card:  card,
		Bid:   bid,
		Combo: combo,
	}
}

var cardRank = map[rune]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 11, 'T': 10,
	'9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

func cardPower(c rune) int {
	return cardRank[c]
}

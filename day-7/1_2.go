package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var fileContent, fileErr = os.ReadFile("input.txt")

func main() {
	fmt.Println(evaluateOne(string(fileContent)))
	fmt.Println(evaluateTwo(string(fileContent)))
}

var handRanks = [7][]int{
	{1, 1, 1, 1, 1},
	{1, 1, 1, 2},
	{1, 2, 2},
	{1, 1, 3},
	{2, 3},
	{1, 4},
	{5},
}

type PokerHand struct {
	cardSet string
	wager   int
	order   int
}

func evaluateOne(inputData string) int {
	inputLines := strings.Split(inputData, "\n")

	var strengthMap = make(map[int][]PokerHand)

	for _, eachLine := range inputLines {
		if len(eachLine) == 0 {
			continue
		}

		cardSequence := strings.Fields(eachLine)[0]
		wagerValue, _ := strconv.Atoi(strings.Fields(eachLine)[1])
		strengthVal := calculateStrength(cardSequence, nil, false, 0)

		strengthMap[strengthVal] = append(strengthMap[strengthVal], PokerHand{cardSequence, wagerValue, 0})
	}

	for _, handsArray := range strengthMap {
		if len(handsArray) > 1 {
			sort.Slice(handsArray, func(i, j int) bool {
				return cardComparison(handsArray[i].cardSet, handsArray[j].cardSet, cardStrengthOne)
			})
		}
	}

	var strengthValues []int
	for strengthKey := range strengthMap {
		strengthValues = append(strengthValues, strengthKey)
	}

	sort.Ints(strengthValues)

	var totalSum int
	currentRank := 1
	for _, strength := range strengthValues {
		for _, hand := range strengthMap[strength] {
			totalSum += hand.wager * currentRank
			currentRank++
		}
	}

	return totalSum
}

func evaluateTwo(inputData string) int {
	inputLines := strings.Split(inputData, "\n")

	var strengthMap = make(map[int][]PokerHand)

	for _, eachLine := range inputLines {
		if len(eachLine) == 0 {
			continue
		}

		cardSequence := strings.Fields(eachLine)[0]
		wagerValue, _ := strconv.Atoi(strings.Fields(eachLine)[1])
		strengthVal := calculateStrength(cardSequence, nil, true, 0)

		strengthMap[strengthVal] = append(strengthMap[strengthVal], PokerHand{cardSequence, wagerValue, 0})
	}

	for _, handsArray := range strengthMap {
		if len(handsArray) > 1 {
			sort.Slice(handsArray, func(i, j int) bool {
				return cardComparison(handsArray[i].cardSet, handsArray[j].cardSet, cardStrengthTwo)
			})
		}
	}

	var strengthValues []int
	for strengthKey := range strengthMap {
		strengthValues = append(strengthValues, strengthKey)
	}

	sort.Ints(strengthValues)

	var totalSum int
	currentRank := 1
	for _, strength := range strengthValues {
		for _, hand := range strengthMap[strength] {
			totalSum += hand.wager * currentRank
			currentRank++
		}
	}

	return totalSum
}

func cardComparison(firstCards, secondCards string, strengthFunc func(byte) int) bool {
	for idx := 0; idx < len(firstCards); idx++ {
		if strengthFunc(firstCards[idx]) == strengthFunc(secondCards[idx]) {
			continue
		}

		return strengthFunc(firstCards[idx]) < strengthFunc(secondCards[idx])
	}

	return false
}

func cardStrengthOne(card byte) int {
	return strings.Index("23456789TJQKA", string(card))
}

func cardStrengthTwo(card byte) int {
	return strings.Index("J23456789TQKA", string(card))
}

func calculateStrength(cardSequence string, pairCount []int, useJokers bool, jokerCount int) int {
	if len(cardSequence) == 0 {
		if useJokers && jokerCount > 0 {
			pairCount = jokerReplace(pairCount, jokerCount)
		}
		sort.Ints(pairCount)
		for i, rankPattern := range handRanks {
			if compareSlices(rankPattern, pairCount) == 0 {
				return i + 1
			}
		}
	}

	matchCount := 1
	cardCompare := string(cardSequence[0])
	if cardCompare == "J" {
		jokerCount++
	}
	cardSequence = cardSequence[1:]

	for strings.Index(cardSequence, cardCompare) != -1 {
		matchCount++
		if cardCompare == "J" {
			jokerCount++
		}
		cardSequence = strings.Replace(cardSequence, cardCompare, "", 1)
	}
	pairCount = append(pairCount, matchCount)

	return calculateStrength(cardSequence, pairCount, useJokers, jokerCount)
}

func jokerReplace(pairArr []int, jokerNum int) []int {
	jokerIdx := sliceIndex(pairArr, jokerNum)
	if jokerIdx == -1 || jokerNum == 5 {
		return pairArr
	}
	pairArr = append(pairArr[:jokerIdx], pairArr[jokerIdx+1:]...)
	sort.Ints(pairArr)
	pairArr[len(pairArr)-1] += jokerNum

	return pairArr
}

func sliceIndex(slice []int, value int) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func compareSlices(slice1, slice2 []int) int {
	if len(slice1) != len(slice2) {
		return -1
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return -1
		}
	}
	return 0
}

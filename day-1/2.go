package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("calibration.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Mapping of number words to their digit values
	numberWords := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	var totalSum int
	scanner := bufio.NewScanner(file)
	// Regex to match whole words or single digits, prioritizing words
	digitRegex := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := digitRegex.FindAllString(line, -1)

		if len(matches) > 0 {
			first := matches[0]
			last := matches[len(matches)-1]

			// Convert words to digits if necessary
			firstDigit, ok := numberWords[first]
			if !ok {
				firstDigit, _ = strconv.Atoi(first)
			}
			lastDigit, ok := numberWords[last]
			if !ok {
				lastDigit, _ = strconv.Atoi(last)
			}

			twoDigitNumber := firstDigit*10 + lastDigit
			totalSum += twoDigitNumber
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("Total sum of calibration values:", totalSum)
}

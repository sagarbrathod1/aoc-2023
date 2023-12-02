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

    var totalSum int
    scanner := bufio.NewScanner(file)
    digitRegex := regexp.MustCompile(`\d`)

    for scanner.Scan() {
        line := scanner.Text()
        digits := digitRegex.FindAllString(line, -1)

        if len(digits) > 0 {
            firstDigit, _ := strconv.Atoi(digits[0])
            lastDigit, _ := strconv.Atoi(digits[len(digits)-1])
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

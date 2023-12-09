package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var histories [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		var history []int
		for _, s := range strings.Split(line, " ") {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			history = append(history, num)
		}
		histories = append(histories, history)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var extrapolatedValues []int

	for _, history := range histories {
		sublists := [][]int{history}
		allZeroes := false

		for !allZeroes {
			var sublist []int
			last := sublists[len(sublists)-1]
			for i := 0; i < len(last)-1; i++ {
				difference := last[i+1] - last[i]
				sublist = append(sublist, difference)
			}

			allZeroes = true
			for _, element := range sublist {
				if element != 0 {
					allZeroes = false
					break
				}
			}

			sublists = append(sublists, sublist)
		}

		sublists[len(sublists)-1] = append(sublists[len(sublists)-1], 0)
		for i := len(sublists) - 2; i >= 0; i-- {
			extrapolatedValue := sublists[i][len(sublists[i])-1] + sublists[i+1][len(sublists[i+1])-1]
			sublists[i] = append(sublists[i], extrapolatedValue)
		}

		extrapolatedValues = append(extrapolatedValues, sublists[0][len(sublists[0])-1])
	}

	var ans int
	for _, val := range extrapolatedValues {
		ans += val
	}

	fmt.Println(ans)
}

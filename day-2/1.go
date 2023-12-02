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
        panic(err)
    }
    defer file.Close()

    var possibleGamesSum int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ": ")
        gameIDStr := parts[0][5:]
        gameID, _ := strconv.Atoi(gameIDStr)

        draws := strings.Split(parts[1], "; ")
        possible := true
        for _, draw := range draws {
            cubes := strings.Split(draw, ", ")
            red, green, blue := 0, 0, 0
            for _, cube := range cubes {
                cubeInfo := strings.Split(cube, " ")
                count, _ := strconv.Atoi(cubeInfo[0])
                color := cubeInfo[1]
                switch color {
                case "red":
                    red += count
                case "green":
                    green += count
                case "blue":
                    blue += count
                }
            }
            if red > 12 || green > 13 || blue > 14 {
                possible = false
                break
            }
        }
        if possible {
            possibleGamesSum += gameID
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Printf("Sum of IDs of possible games: %d\n", possibleGamesSum)
}


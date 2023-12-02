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

    var totalPower int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ": ")
        
        draws := strings.Split(parts[1], "; ")
        maxRed, maxGreen, maxBlue := 0, 0, 0
        for _, draw := range draws {
            red, green, blue := 0, 0, 0
            cubes := strings.Split(draw, ", ")
            for _, cube := range cubes {
                cubeInfo := strings.Split(cube, " ")
                count, _ := strconv.Atoi(cubeInfo[0])
                color := cubeInfo[1]
                switch color {
                case "red":
                    red = count
                case "green":
                    green = count
                case "blue":
                    blue = count
                }
                if red > maxRed {
                    maxRed = red
                }
                if green > maxGreen {
                    maxGreen = green
                }
                if blue > maxBlue {
                    maxBlue = blue
                }
            }
        }
        power := maxRed * maxGreen * maxBlue
        totalPower += power
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    fmt.Printf("Sum of the power of the minimum sets: %d\n", totalPower)
}


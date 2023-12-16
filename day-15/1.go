package main

import (
	"fmt"
	"os"
	"strings"
)

func hash(s string) int {
    currentValue := 0
    for _, char := range s {
        asciiValue := int(char)
        currentValue += asciiValue
        currentValue *= 17
        currentValue %= 256
    }
    return currentValue
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run script.go <filename>")
        os.Exit(1)
    }

    filename := os.Args[1]
    data, err := os.ReadFile(filename)
    if err != nil {
        fmt.Println("Error reading file:", err)
        os.Exit(1)
    }

    sequence := string(data)
    sequence = strings.TrimSpace(sequence)
    steps := strings.Split(sequence, ",")
    sum := 0
    for _, step := range steps {
        sum += hash(step)
    }
    fmt.Println(sum)
}

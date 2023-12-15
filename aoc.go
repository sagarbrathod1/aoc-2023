package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	baseURL = "https://adventofcode.com"
	sessionCookie = "53616c7465645f5ff2d0c51ca6ea23c416be5958f715c4bf5d8c6e0c09986f9a8c5ce269ca18e76b81f2c0bf1a7811b40bac9cd460a6e88a8a83c4349a6a0ad3"
	baseFolder = "." 
)

func getPuzzleInput(day int) (io.ReadCloser, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/2023/day/%d/input", baseURL, day), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Unexpected Status Code: %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func savePuzzleInput(day int, input io.ReadCloser) error {
	defer input.Close()

	folderName, err := ensureFolderExists(day)
	if err != nil {
			return err
	}

	filePath := fmt.Sprintf("%s/puzzle.txt", folderName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, input)
	return err
}

func ensureFolderExists(day int) (string, error) {
	folderName := fmt.Sprintf("day-%d", day)
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
			return "", err
	}
	return folderName, nil
}

func createBoilerplate(day int) error {
	folderName, err := ensureFolderExists(day)
	if err != nil {
			return err
	}

	for i := 1; i <= 2; i++ {
		fileName := fmt.Sprintf("%s/%d.go", folderName, i)
		if err := writeBoilerplate(fileName); err != nil {
			return err
		}
	}

	return nil
}

func writeBoilerplate(filePath string) error {
	if _, err := os.Stat(filePath); err == nil {
		return fmt.Errorf("File already exists: %s", filePath)
	} else if !os.IsNotExist(err) {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	boilerplate := 
	`
	package main

	import (
		"bufio"
		"fmt"
		"os"
	)

	func main() {
		file, err := os.Open("puzzle.txt")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close()

		total := 0

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			// Process line
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading from file:", err)
			return
		}

		fmt.Println("Part 1 Solution:", total)
	}
	`

	_, err = file.WriteString(boilerplate)
	return err
}

func main() {
	day := time.Now().Day() 

	input, err := getPuzzleInput(day)
	if err != nil {
		fmt.Printf("Error getting input: %s\n", err)
		return
	}

	if err := savePuzzleInput(day, input); err != nil {
		fmt.Printf("Error saving input: %s\n", err)
		return
	}
	fmt.Printf("Day %d input saved successfully.\n", day)

	if err := createBoilerplate(day); err != nil {
		fmt.Printf("Error creating boilerplate: %s\n", err)
		return
	}
	fmt.Println("Boilerplate created successfully.")
}


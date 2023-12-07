package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MapEntry struct {
	destStart, sourceStart, rangeLength int64
}

func getSeeds(scanner *bufio.Scanner) [][]int64 {
	var seeds [][]int64
	scanner.Scan() 
	seedsLine := scanner.Text()
	seedsParts := strings.Fields(seedsLine)[1:]
	for i := 0; i < len(seedsParts); i += 2 {
		start, _ := strconv.ParseInt(seedsParts[i], 10, 64)
		length, _ := strconv.ParseInt(seedsParts[i+1], 10, 64)
		seeds = append(seeds, []int64{start, length})
	}
	return seeds
}

func getMap(scanner *bufio.Scanner) []MapEntry {
	var mapping []MapEntry
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Fields(line)
		if len(parts) == 3 {
			destStart, _ := strconv.ParseInt(parts[0], 10, 64)
			sourceStart, _ := strconv.ParseInt(parts[1], 10, 64)
			rangeLength, _ := strconv.ParseInt(parts[2], 10, 64)
			mapping = append(mapping, MapEntry{destStart: destStart, sourceStart: sourceStart, rangeLength: rangeLength})
		}
	}
	return mapping
}

func processMap(seed int64, mapping []MapEntry) int64 {
	for _, entry := range mapping {
		if seed >= entry.sourceStart && seed < entry.sourceStart+entry.rangeLength {
			return entry.destStart + (seed - entry.sourceStart)
		}
	}
	return seed
}

func main() {
	filePath := "puzzle.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seedRanges := getSeeds(scanner)
	fmt.Println("Seed Ranges:", seedRanges)

	scanner.Scan() 
	seedToSoilMap := getMap(scanner)
	scanner.Scan() 
	soilToFertilizerMap := getMap(scanner)
	scanner.Scan() 
	fertilizerToWaterMap := getMap(scanner)
	scanner.Scan() 
	waterToLightMap := getMap(scanner)
	scanner.Scan() 
	lightToTemperatureMap := getMap(scanner)
	scanner.Scan() 
	temperatureToHumidityMap := getMap(scanner)
	scanner.Scan() 
	humidityToLocationMap := getMap(scanner)

	var lowestLocation int64 = int64(^uint64(0) >> 1) // max int64

	for _, seedRange := range seedRanges {
		for i := int64(0); i < seedRange[1]; i++ {
			currentSeed := seedRange[0] + i
			currentSeed = processMap(currentSeed, seedToSoilMap)
			currentSeed = processMap(currentSeed, soilToFertilizerMap)
			currentSeed = processMap(currentSeed, fertilizerToWaterMap)
			currentSeed = processMap(currentSeed, waterToLightMap)
			currentSeed = processMap(currentSeed, lightToTemperatureMap)
			currentSeed = processMap(currentSeed, temperatureToHumidityMap)
			currentSeed = processMap(currentSeed, humidityToLocationMap)
			if currentSeed < lowestLocation {
				lowestLocation = currentSeed
			}
		}
	}

	fmt.Printf("The lowest location number is: %d\n", lowestLocation)
}

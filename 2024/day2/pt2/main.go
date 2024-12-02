package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Advent of Code 2024 - Day 2 - Part 2
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	// Read day1/input.txt & split it by new line
	lines, err := ReadLines("../../input/day2.txt")
	if err != nil {
		panic(err)
	}

	// Identify safe lines
	safeLineCount := 0

	// Iterate over the lines
	for _, line := range lines {
		// split the line by space and convert the strings to integers
		chunks := strings.Split(line, " ")

		// First check if the full line is valid
		isValid := lineIsValid(chunks)

		if isValid {
			safeLineCount++
		} else {
			// Now check if the line is valid when excluding a chunk
			for i := 0; i < len(chunks); i++ {
				// Create a new slice with the chunk removed
				var newChunks []string
				for j := 0; j < len(chunks); j++ {
					if j != i {
						newChunks = append(newChunks, chunks[j])
					}
				}

				// Check if the new line is valid
				if lineIsValid(newChunks) {
					safeLineCount++
					break
				}
			}
		}
	}

	fmt.Println("Safe Line Count: ", safeLineCount)
}

func lineIsValid(chunks []string) bool {
	positiveDifference := false
	for i := 0; i < len(chunks); i++ {
		leftNum, err := strconv.Atoi(chunks[i])
		if err != nil {
			panic(err)
		}

		if i+1 < len(chunks) {
			rightNum, err := strconv.Atoi(chunks[i+1])
			if err != nil {
				panic(err)
			}

			if i == 0 {

				if rightNum-leftNum > 0 {
					positiveDifference = true
				}
			} else {

				if rightNum-leftNum < 0 && positiveDifference {
					return false
				}

				if rightNum-leftNum > 0 && !positiveDifference {
					return false
				}
			}

			if leftNum == rightNum || rightNum-leftNum > 3 || rightNum-leftNum < -3 {
				return false
			}
		}
	}
	return true
}

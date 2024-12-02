package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Advent of Code 2024 - Day 1 - Part 2
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
	lines, err := ReadLines("../input/day2.txt")
	if err != nil {
		panic(err)
	}

	// Identify safe lines
	safeLineCount := 0

	// Iterate over the lines
	for _, line := range lines {
		// split the line by space and convert the strings to integers
		chunks := strings.Split(line, " ")

		// Compare chunks left to right, if the difference is ever greater than 3 break, or if the difference changes from positive to negative break
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
					// Set positiveDifference to true if the first difference is positive
					if rightNum-leftNum > 0 {
						positiveDifference = true
					}
				} else {
					// Break if the difference changes from positive to negative
					if rightNum-leftNum < 0 && positiveDifference {
						break
					}

					// Break if the difference changes from negative to positive
					if rightNum-leftNum > 0 && !positiveDifference {
						break
					}
				}

				if leftNum == rightNum || rightNum-leftNum > 3 || rightNum-leftNum < -3 {
					break
				}
			}

			if i == len(chunks)-1 {
				safeLineCount++
			}
		}
	}

	fmt.Println("Safe Line Count: ", safeLineCount)
}

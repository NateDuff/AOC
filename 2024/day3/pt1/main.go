package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Advent of Code 2024 - Day 2 - Part 1
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
	lines, err := ReadLines("../../input/day3.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	// Iterate over the lines
	for _, line := range lines {
		// Parse out all instances of mul(386,104) where the numbers are separated by a comma and could be 1 to 3 digits long
		// This is a very simple regex pattern that will match the above example
		regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		// Find all matches in the line
		matches := regex.FindAllStringSubmatch(line, -1)

		// Print out all matches
		for _, match := range matches {
			fmt.Println(match)

			// Convert the strings to integers
			leftNum, err := strconv.Atoi(match[1])
			if err != nil {
				panic(err)
			}

			rightNum, err := strconv.Atoi(match[2])
			if err != nil {
				panic(err)
			}

			total += leftNum * rightNum
		}
	}

	fmt.Println(total)
}

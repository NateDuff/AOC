package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Advent of Code 2024 - Day 2 - Part 1
func ReadAllLines(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content, scanner.Err()
}

func main() {
	// Read day1/input.txt & split it by new line
	content, err := ReadAllLines("../../input/day3.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	// Parse out all instances of mul(386,104) where the numbers are separated by a comma and could be 1 to 3 digits long, or do() or don't()
	// This is a very simple regex pattern that will match the above example
	//regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	regex := regexp.MustCompile(`(do\(\)|don't\(\)|mul\((\d{1,3}),(\d{1,3})\))`)

	// Find all matches in the line
	matches := regex.FindAllStringSubmatch(content, -1)

	enabled := true
	// Print out all matches
	for _, match := range matches {
		fmt.Println(match)

		if match[1] == "do()" {
			enabled = true
			continue
		}

		if match[1] == "don't()" {
			enabled = false
			continue
		}

		if !enabled {
			continue
		}

		// Convert the strings to integers
		leftNum, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		rightNum, err := strconv.Atoi(match[3])
		if err != nil {
			panic(err)
		}

		total += leftNum * rightNum
	}

	fmt.Println(total)
}

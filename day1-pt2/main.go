package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	lines, err := ReadLines("../input/day1.txt")
	if err != nil {
		panic(err)
	}

	// Each line contains 2 numbers separated by a space
	leftGroup := make([]int, 0)
	rightGroup := make([]int, 0)

	// Iterate over the lines
	for _, line := range lines {
		// split the line by space and convert the strings to integers
		chunks := strings.Split(line, " ")
		leftNum, err := strconv.Atoi(chunks[0])
		if err != nil {
			panic(err)
		}
		leftGroup = append(leftGroup, leftNum)

		rightNum, err := strconv.Atoi(chunks[3])
		if err != nil {
			panic(err)
		}
		rightGroup = append(rightGroup, rightNum)
	}

	// Sort the groups
	sort.Ints(leftGroup)
	sort.Ints(rightGroup)

	// Count the number of lines in each group
	leftGroupCount := len(leftGroup)
	fmt.Println("Left Group Count: ", leftGroupCount)

	rightGroupCount := len(rightGroup)
	fmt.Println("Right Group Count: ", rightGroupCount)

	similarityScoreMap := make(map[int]int)
	similarityScoreGroup := make([]int, 0)

	// Calculate the similarity score for each unique number in the left group with the right group
	for _, leftNum := range leftGroup {
		// Check if the leftNum is already in the similarityScoreMap
		if _, ok := similarityScoreMap[leftNum]; ok {
			continue
		}

		// Count the number of times the leftNum appears in the right group
		rightNumCount := 0
		for _, num := range rightGroup {
			if leftNum == num {
				rightNumCount++
			}
		}

		// Calculate the similarity score
		similarityScore := rightNumCount * leftNum
		similarityScoreMap[leftNum] = similarityScore
	}

	// Loop through leftGroup and get the similarity score for each number
	for _, num := range leftGroup {
		similarityScoreGroup = append(similarityScoreGroup, similarityScoreMap[num])
	}

	// Sum the similarityScore group
	sum := 0
	for _, score := range similarityScoreGroup {
		sum += score
	}

	fmt.Println("Sum: ", sum)
}

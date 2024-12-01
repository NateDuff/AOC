package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Advent of Code 2024 - Day 1 - Part 1
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

	differenceGroup := make([]int, 0)
	// Loop through groups, compare left and right for each index and set the difference group
	for i := 0; i < leftGroupCount; i++ {
		left := leftGroup[i]
		right := rightGroup[i]

		diff := right - left

		if diff < 0 {
			diff = diff * -1
		}

		differenceGroup = append(differenceGroup, diff)
	}

	// Sum the difference group
	sum := 0
	for _, diff := range differenceGroup {
		sum += diff
	}

	fmt.Println("Sum: ", sum)
}

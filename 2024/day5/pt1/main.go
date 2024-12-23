package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Advent of Code 2024 - Day 5 - Part 1
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

func checkRules(line string, rules map[int][]int) bool {
	// Rules consist of two numbers, check if the line contains the first number, if it does, check if it contains the second number and return false when the second number comes after the first number
	lineIsValid := true

	for firstNum, secondNum := range rules {
		if strings.Contains(line, strconv.Itoa(firstNum)) {
			for _, num := range secondNum {
				if strings.Contains(line, strconv.Itoa(num)) {

					// Check if the second number comes after the first number
					firstIndex := strings.Index(line, strconv.Itoa(firstNum))
					secondIndex := strings.Index(line, strconv.Itoa(num))

					if firstIndex > secondIndex {
						lineIsValid = false
					}
				}
			}
		}
	}

	return lineIsValid
}

func getCenterNumber(line string) int {
	parts := strings.Split(line, ",")
	partCount := len(parts)
	middleIndex := partCount / 2

	centerNum, err := strconv.Atoi(parts[middleIndex])
	if err != nil {
		panic(err)
	}

	return centerNum
}

func main() {
	// Part 1 - get rules
	rules, err := ReadLines("../../input/day5-a.txt")
	if err != nil {
		panic(err)
	}

	ruleSets := make(map[int][]int)

	for _, rule := range rules {
		// split on | to get the two parts
		parts := strings.Split(rule, "|")

		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		secondNum, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		ruleSets[firstNum] = append(ruleSets[firstNum], secondNum)
	}

	// Part 2 - check if input lines match the rules
	inputLines, err := ReadLines("../../input/day5-b.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	for _, line := range inputLines {
		if checkRules(line, ruleSets) {
			total += getCenterNumber(line)
		} else {
			fmt.Println("Line does not match the rules:", line)
		}
	}

	fmt.Println(total)
}

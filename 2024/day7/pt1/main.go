package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Advent of Code 2024 - Day 7 - Part 1
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

func addMul(result int, r1 int, depth int, tokens []string) bool {
	// Check if we've exceeded our depth
	if depth >= len(tokens) {
		return false
	}

	r2, err := strconv.Atoi(tokens[depth])
	if err != nil {
		return false
	}

	// Multiplication check
	if (r1 * r2) == result {
		return true
	} else {
		if addMul(result, r1*r2, depth+1, tokens) {
			return true
		}
	}

	// Addition check
	if (r1 + r2) == result {
		return true
	} else {
		if addMul(result, r1+r2, depth+1, tokens) {
			return true
		}
	}

	return false
}

func main() {
	lines, err := ReadLines("../../input/day7.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		subTotal, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		subParts := strings.Split(parts[1], " ")

		firstSubPart, _ := strconv.Atoi(subParts[0])
		if addMul(subTotal, firstSubPart, 1, subParts) {
			total += subTotal
		}
	}

	fmt.Println(total)
}

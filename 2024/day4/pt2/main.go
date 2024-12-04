package main

import (
	"bufio"
	"fmt"
	"os"
)

// Advent of Code 2024 - Day 4 - Part 2
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

func findWord(board []string) int {
	total := 0
	rows := len(board)
	cols := len(board[0])

	// Iterate through inner cells
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			// Check for X-MAS pattern for each 'A' cell
			if board[r][c] == 'A' && checkConditions(board, r, c) {
				total++
			}
		}
	}

	return total
}

func checkConditions(board []string, r, c int) bool {
	// M - S
	// - A -
	// M - S
	if board[r-1][c-1] == 'M' && // Top-left
		board[r+1][c+1] == 'S' && // Bottom-right
		board[r-1][c+1] == 'S' && // Top-right
		board[r+1][c-1] == 'M' { // Bottom-left
		return true
	}

	// S - M
	// - A -
	// S - M
	if board[r-1][c-1] == 'S' && // Top-left
		board[r+1][c+1] == 'M' && // Bottom-right
		board[r-1][c+1] == 'M' && // Top-right
		board[r+1][c-1] == 'S' { // Bottom-left
		return true
	}

	// S - S
	// - A -
	// M - M
	if board[r-1][c-1] == 'S' && // Top-left
		board[r+1][c+1] == 'M' && // Bottom-right
		board[r-1][c+1] == 'S' && // Top-right
		board[r+1][c-1] == 'M' { // Bottom-left
		return true
	}

	// M - M
	// - A -
	// S - S
	if board[r-1][c-1] == 'M' && // Top-left
		board[r+1][c+1] == 'S' && // Bottom-right
		board[r-1][c+1] == 'M' && // Top-right
		board[r+1][c-1] == 'S' { // Bottom-left
		return true
	}

	return false
}

func main() {
	lines, err := ReadLines("../../input/day4.txt")
	if err != nil {
		panic(err)
	}

	total := findWord(lines)

	fmt.Println(total)
}

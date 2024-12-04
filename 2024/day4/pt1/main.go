package main

import (
	"bufio"
	"fmt"
	"os"
)

// Advent of Code 2024 - Day 4 - Part 1
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

func findWord(board []string, word string) int {
	total := 0
	rows := len(board)
	cols := len(board[0])

	// Directions: right, left, down, up, diagonals
	directions := [][]int{
		{0, 1},   // Right
		{0, -1},  // Left
		{1, 0},   // Down
		{-1, 0},  // Up
		{1, 1},   // Diagonal Down-Right
		{-1, -1}, // Diagonal Up-Left
		{1, -1},  // Diagonal Down-Left
		{-1, 1},  // Diagonal Up-Right
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				if searchFrom(board, word, r, c, dir[0], dir[1]) {
					fmt.Printf("Word found starting at row %d, col %d in direction %+v\n", r, c, dir)
					total++
				}
			}
		}
	}
	return total
}

func searchFrom(board []string, word string, startRow, startCol, dRow, dCol int) bool {
	rows := len(board)
	cols := len(board[0])
	wordLen := len(word)

	for i := 0; i < wordLen; i++ {
		newRow := startRow + i*dRow
		newCol := startCol + i*dCol

		// Check bounds
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
			return false
		}

		// Check character match
		if board[newRow][newCol] != word[i] {
			return false
		}
	}

	return true
}

func main() {
	lines, err := ReadLines("../../input/day4.txt")
	if err != nil {
		panic(err)
	}

	total := findWord(lines, "XMAS")

	fmt.Println(total)
}

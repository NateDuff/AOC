package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

type index struct {
	r, c int
}

func parseInput(fileName string) ([][]rune, index) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	var startIndex index
	for scanner.Scan() {
		row := []rune(scanner.Text())
		grid = append(grid, row)
		if c := slices.Index(row, '^'); c != -1 {
			startIndex = index{r: len(grid) - 1, c: c}
			grid[startIndex.r][startIndex.c] = '.'
		}
	}

	return grid, startIndex
}

func traverseGrid(grid [][]rune, startIndex index) ([]index, int) {
	directions := []index{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	maxRow, maxCol := len(grid), len(grid[0])
	visited := make(map[index]bool)

	at, facing := startIndex, 0

	for isValidIndex(at, maxRow, maxCol) {
		if grid[at.r][at.c] == '#' {
			at.r -= directions[facing].r
			at.c -= directions[facing].c
			facing = (facing + 1) % len(directions)
			continue
		}
		visited[at] = true
		grid[at.r][at.c] = 'X'
		at.r += directions[facing].r
		at.c += directions[facing].c
	}

	visitedIndices := make([]index, 0, len(visited))
	for idx := range visited {
		visitedIndices = append(visitedIndices, idx)
	}

	return visitedIndices, len(visitedIndices)
}

func countCycles(grid [][]rune, startIndex index, visitedIndices []index) int {
	cycleCount := 0
	for _, idx := range visitedIndices {
		if idx == startIndex || grid[idx.r][idx.c] == '#' {
			continue
		}
		grid[idx.r][idx.c] = '#'
		if hasCycle(grid, startIndex) {
			cycleCount++
		}
		grid[idx.r][idx.c] = '.'
	}

	return cycleCount
}

func hasCycle(grid [][]rune, startIndex index) bool {
	visited := make(map[index]index)
	directions := []index{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	maxRow, maxCol := len(grid), len(grid[0])
	at, facing := startIndex, 0

	for isValidIndex(at, maxRow, maxCol) {
		if visited[at] == directions[facing] {
			return true
		}
		visited[at] = directions[facing]

		if grid[at.r][at.c] == '#' {
			at.r -= directions[facing].r
			at.c -= directions[facing].c
			facing = (facing + 1) % len(directions)
			continue
		}

		grid[at.r][at.c] = 'X'
		at.r += directions[facing].r
		at.c += directions[facing].c
	}

	return false
}

func isValidIndex(idx index, maxRow, maxCol int) bool {
	return idx.r >= 0 && idx.c >= 0 && idx.r < maxRow && idx.c < maxCol
}

func main() {
	puzzleInput, startIndex := parseInput("../../input/day6.txt")
	visitedIndices, _ := traverseGrid(puzzleInput, startIndex)

	fmt.Println(countCycles(puzzleInput, startIndex, visitedIndices))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Position struct {
	x, y int
}

const (
	East     = 0
	South    = 1
	West     = 2
	North    = 3
	Free     = 4
	Obstacle = 5
)

var directions = []Position{
	East:  {1, 0},  // East
	South: {0, 1},  // South
	West:  {-1, 0}, // West
	North: {0, -1}, // North
}

// isMoveValid checks if a move from currentPos to targetPos is valid.
func isMoveValid(grid [][]int, currentPos, targetPos Position) bool {
	if targetPos.x < 0 || targetPos.x >= len(grid[0]) || targetPos.y < 0 || targetPos.y >= len(grid) {
		return false
	}
	if grid[targetPos.y][targetPos.x]-grid[currentPos.y][currentPos.x] != 1 {
		return false
	}
	return true
}

// solveDfs solves the maze using depth-first search.
func solveDfs(grid [][]int, currentPos, endingPos Position, visited map[Position]bool, currentPath []Position) ([]Position, int) {
	visited[currentPos] = true
	currentPath = append(currentPath, currentPos)
	var longestPath []Position
	maxLength := 0

	if currentPos == endingPos {
		if len(currentPath) > len(longestPath) {
			longestPath = make([]Position, len(currentPath))
			copy(longestPath, currentPath)
		}
		return longestPath, 1
	}

	for _, direction := range directions {
		nextPos := Position{currentPos.x + direction.x, currentPos.y + direction.y}
		if isMoveValid(grid, currentPos, nextPos) {
			newPath, length := solveDfs(grid, nextPos, endingPos, visited, currentPath)

			if length > maxLength {
				maxLength = length
				longestPath = make([]Position, len(newPath))
				copy(longestPath, newPath)
			}
		}
	}

	return longestPath, maxLength
}

// readGridWithPositions reads a grid from a file and returns the grid, the starting points and the ending points.
func readGridWithPositions(filename string) ([][]int, []Position, []Position) {
	var grid [][]int
	var startingPoints, endingPoints []Position

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineY := 0
	for scanner.Scan() {
		var line []int
		lineText := scanner.Text()
		for lineX, char := range lineText {
			num, _ := strconv.Atoi(string(char))
			if num == 0 {
				startingPoints = append(startingPoints, Position{lineX, lineY})
			} else if num == 9 {
				endingPoints = append(endingPoints, Position{lineX, lineY})
			}
			line = append(line, num)
		}
		grid = append(grid, line)
		lineY++
	}
	return grid, startingPoints, endingPoints
}

// longestPath returns the longest path from start to end in the grid.
func longestPath(grid [][]int, start, end Position) int {
	visited := make(map[Position]bool)

	_, length := solveDfs(grid, start, end, visited, []Position{})

	return length
}

func main() {
	grid, startingPoints, endingPoints := readGridWithPositions("../../input/day10.txt")
	total := 0

	for _, start := range startingPoints {
		for _, end := range endingPoints {
			total += longestPath(grid, start, end)
		}
	}

	fmt.Println(total)
}

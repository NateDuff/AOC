package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func parseInput(input string) map[image.Point]rune {
	grid := map[image.Point]rune{}
	for y, s := range strings.Fields(input) {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}
	return grid
}

func calculateAreaAndPerimeter(grid map[image.Point]rune, seen map[image.Point]bool, start image.Point) (int, int, int) {
	area := 1
	perimeter, sides := 0, 0
	queue := []image.Point{start}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			n := p.Add(d)
			if grid[n] != grid[p] {
				rotatedPoint := p.Add(image.Point{-d.Y, d.X})
				if grid[rotatedPoint] != grid[p] || grid[rotatedPoint.Add(d)] == grid[p] {
					sides++
				}
				perimeter++
			} else if !seen[n] {
				seen[n] = true
				queue = append(queue, n)
				area++
			}
		}
	}
	return area, perimeter, sides
}

func main() {
	input, _ := os.ReadFile("../../input/day12.txt")

	grid := parseInput(string(input))

	seen := map[image.Point]bool{}
	part1, part2 := 0, 0
	for p := range grid {
		if seen[p] {
			continue
		}
		seen[p] = true

		area, perimeter, sides := calculateAreaAndPerimeter(grid, seen, p)
		part1 += area * perimeter
		part2 += area * sides
	}
	fmt.Println("Part 1:", part1, "Part 2:", part2)
}

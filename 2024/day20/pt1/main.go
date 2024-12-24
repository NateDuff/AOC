package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

func absoluteValue(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input, _ := os.ReadFile("../../input/day20.txt")

	grid, start := map[image.Point]rune{}, image.Point{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			if r == 'S' {
				start = image.Point{x, y}
			}
			grid[image.Point{x, y}] = r
		}
	}

	queue, dist := []image.Point{start}, map[image.Point]int{start: 0}
	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		for _, d := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
			n := p.Add(d)
			if _, ok := dist[n]; !ok && grid[n] != '#' {
				queue, dist[n] = append(queue, n), dist[p]+1
			}
		}
	}

	part1, part2 := 0, 0
	for p1 := range dist {
		for p2 := range dist {
			d := absoluteValue(p2.X-p1.X) + absoluteValue(p2.Y-p1.Y)
			if d <= 20 && dist[p2] >= dist[p1]+d+100 {
				if d <= 2 {
					part1++
				}
				part2++
			}
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func countWays(design string, pats []string, cache map[string]int) int {
	if val, ok := cache[design]; ok {
		return val
	}
	if design == "" {
		return 1
	}

	ways := 0
	for _, pat := range pats {
		if strings.HasPrefix(design, pat) {
			ways += countWays(design[len(pat):], pats, cache)
		}
	}
	cache[design] = ways
	return ways
}

func main() {
	input, _ := os.ReadFile("../../input/day19.txt")

	split := strings.Split(string(input), "\r\n\r\n")

	availablePatterns := strings.Split(strings.TrimSpace(split[0]), ", ")
	desiredPatterns := strings.Fields(split[1])

	cache := make(map[string]int)

	part1, part2 := 0, 0
	for _, pattern := range desiredPatterns {
		if ways := countWays(pattern, availablePatterns, cache); ways > 0 {
			part1++
			part2 += ways
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}

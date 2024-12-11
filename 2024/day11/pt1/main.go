package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// blinkNTimes performs the blink operation iteratively for the given number of iterations.
func blinkNTimes(iterations int) int {
	lines := readLinesToList()
	if len(lines) == 0 || len(lines[0]) == 0 {
		return 0
	}

	stones := make(map[int]int)
	for _, stone := range lines[0] {
		stones[stone]++
	}

	for i := 0; i < iterations; i++ {
		newStones := make(map[int]int, len(stones))
		for rock, count := range stones {
			blinkResults := blink(rock)
			for _, blinkResult := range blinkResults {
				newStones[blinkResult] += count
			}
		}
		for k := range stones {
			delete(stones, k)
		}
		for k, v := range newStones {
			stones[k] = v
		}
	}

	total := 0
	for _, count := range stones {
		total += count
	}
	return total
}

// readLinesToList reads integers from a file into a slice of slices of integers.
func readLinesToList() [][]int {
	file, err := os.Open("../../input/day11.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		strValues := strings.Fields(line)
		intValues := make([]int, len(strValues))
		for i, val := range strValues {
			intValues[i], err = strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
		}
		lines = append(lines, intValues)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

// blink implements the blink logic for a given stone.
func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	s := strconv.Itoa(stone)
	if len(s)%2 == 0 {
		mid := len(s) / 2
		part1, _ := strconv.Atoi(s[:mid])
		part2, _ := strconv.Atoi(s[mid:])
		return []int{part1, part2}
	}

	return []int{stone * 2024}
}

func main() {
	fmt.Printf("Blinking %d times: %d\n", 25, blinkNTimes(25))
	fmt.Printf("Blinking %d times: %d\n", 75, blinkNTimes(75))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func fixLineUntilCorrect(line string, rules []string) string {
	fixedLine := line
	previousLine := ""

	for fixedLine != previousLine {
		previousLine = fixedLine
		fixedLine = fixLine(fixedLine, rules)
	}

	return fixedLine
}

func fixLineUntilCheckRulesPasses(line string, rules []string, ruleSets map[int][]int) string {
	fixedLine := line
	previousLine := ""

	for !checkRules(fixedLine, ruleSets) {
		previousLine = fixedLine
		fixedLine = fixLine(fixedLine, rules)
		if fixedLine == previousLine {
			break
		}
	}

	return fixedLine
}

func fixLine(line string, rules []string) string {
	chunks := strings.Split(line, ",")
	for _, rule := range rules {
		ruleParts := strings.Split(rule, "|")
		firstNum, err := strconv.Atoi(ruleParts[0])
		if err != nil {
			panic(err)
		}

		secondNum, err := strconv.Atoi(ruleParts[1])
		if err != nil {
			panic(err)
		}

		firstIndex := findIndex(chunks, firstNum)
		if firstIndex == -1 {
			continue
		}
		chunks = swapChunks(chunks, firstIndex, secondNum)
	}
	return strings.Join(chunks, ",")
}

func findIndex(chunks []string, num int) int {
	for i, chunk := range chunks {
		if chunk == strconv.Itoa(num) {
			return i
		}
	}
	return -1
}

func swapChunks(chunks []string, firstIndex int, secondNum int) []string {
	for i := range chunks {
		if i > firstIndex && chunks[i] == strconv.Itoa(secondNum) {
			chunks[firstIndex], chunks[i] = chunks[i], chunks[firstIndex]
		}
	}

	return chunks
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

func less(nums []int) func(i, j int) bool {
	return func(i, j int) bool {
		return nums[i] > nums[j]
	}
}

func contains(nums []int, num int) bool {
	for _, n := range nums {
		if n == num {
			return true
		}
	}

	return false
}

func getOrderedRules(rules []string) []string {
	firstNumsList := make([]int, 0)

	for _, rule := range rules {
		parts := strings.Split(rule, "|")

		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		// secondNum, err := strconv.Atoi(parts[1])
		// if err != nil {
		// 	panic(err)
		// }

		if !contains(firstNumsList, firstNum) {
			firstNumsList = append(firstNumsList, firstNum)
		}
	}

	orderedRules := make([]string, 0)

	sort.Slice(firstNumsList, less(firstNumsList))

	for _, num := range firstNumsList {
		for _, rule := range rules {
			parts := strings.Split(rule, "|")

			firstNum, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			secondNum, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			if firstNum == num {
				orderedRules = append(orderedRules, strconv.Itoa(firstNum)+"|"+strconv.Itoa(secondNum))
			}
		}
	}

	return orderedRules
}

func getOrderedRuleSets(rules []string, ruleSets map[int][]int) map[int][]int {
	firstNumsList := make([]int, 0)

	for _, rule := range rules {
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
		if !contains(firstNumsList, firstNum) {
			firstNumsList = append(firstNumsList, firstNum)
		}
	}

	orderedRuleSets := make(map[int][]int)

	sort.Slice(firstNumsList, less(firstNumsList))

	for _, firstNum := range firstNumsList {
		orderedRuleSets[firstNum] = ruleSets[firstNum]
	}
	return orderedRuleSets
}

func main() {
	// Part 1 - get rules
	rules, err := ReadLines("../../input/day5-a.txt")
	if err != nil {
		panic(err)
	}

	ruleSets := make(map[int][]int)

	orderedRules := getOrderedRules(rules)
	orderedRuleSets := getOrderedRuleSets(rules, ruleSets)

	// Part 2 - check if input lines match the rules
	inputLines, err := ReadLines("../../input/day5-b.txt")
	if err != nil {
		panic(err)
	}

	total := 0

	for _, line := range inputLines {
		if !checkRules(line, orderedRuleSets) {
			fixedLine := fixLineUntilCheckRulesPasses(line, orderedRules, orderedRuleSets)
			total += getCenterNumber(fixedLine)
		}
	}

	fmt.Println(total)
}

// 5639 too low
// 5665 ... no
// 6098 ... no
// 6101 ... no
// 6336
// 6422 too high

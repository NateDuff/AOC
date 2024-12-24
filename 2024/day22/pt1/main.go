package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadFileLineByLine(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var output []string

	scanner := bufio.NewScanner(file)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return output
}

func FetchNumFromStringIgnoringNonNumeric(line string) int {
	var build strings.Builder
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		return int(localNum)
	}
	return 0
}

func findSecretNumber(input int64) int64 {
	input = prune(mix(int64(64)*input, input))
	input = prune(mix(input/32, input))
	input = prune(mix(input*2048, input))
	return input
}

func mix(value, secretNum int64) int64 {
	return value ^ secretNum
}

func prune(value int64) int64 {
	return value % 16777216
}

type bananaPrice struct {
	num    int
	change int
}

func getNumAndChange(input int64, previous int) bananaPrice {
	return bananaPrice{num: int(input % 10), change: int(input%10) - previous}
}

func calculate(input []string) int64 {
	var sum int64
	b := make([][]bananaPrice, len(input))
	for i, line := range input {
		b[i] = make([]bananaPrice, 2000)
		num := int64(FetchNumFromStringIgnoringNonNumeric(line))
		prev := int(num % 10)
		for j := range 2000 {
			num = findSecretNumber(num)
			b[i][j] = getNumAndChange(num, prev)
			prev = int(num % 10)
		}
		sum += num
	}

	return sum
}

type seq struct {
	a, b, c, d int
}

func main() {
	input := ReadFileLineByLine("../../input/day22.txt")

	fmt.Println(calculate(input))
}

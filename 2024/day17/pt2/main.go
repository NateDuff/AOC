package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func runRecursive(a, b, c int, pgm []int) (out []int) {
	for ip := 0; ip < len(pgm); ip += 2 {
		// Get the current operation and iteration literal
		op, iter := pgm[ip], pgm[ip+1]

		// Get the current value to use
		curr := iter
		switch curr {
		case 4:
			curr = a
		case 5:
			curr = b
		case 6:
			curr = c
		}

		// Execute the operation
		switch op {
		case 0:
			a >>= curr
		case 1:
			b ^= iter
		case 2:
			b = curr % 8
		case 3:
			if a != 0 {
				ip = iter - 2
			}
		case 4:
			b ^= c
		case 5:
			out = append(out, curr%8)
		case 6:
			b = a >> curr
		case 7:
			c = a >> curr
		}
	}
	return out
}

func main() {
	input, _ := os.ReadFile("../../input/day17.txt")

	m := regexp.MustCompile(`[\d,]+`).FindAllString(string(input), -1)

	a, _ := strconv.Atoi(m[0])
	b, _ := strconv.Atoi(m[1])
	c, _ := strconv.Atoi(m[2])
	var pgm []int
	json.Unmarshal([]byte("["+m[3]+"]"), &pgm)

	a = 0
	for n := len(pgm) - 1; n >= 0; n-- {
		a <<= 3
		for !slices.Equal(runRecursive(a, b, c, pgm), pgm[n:]) {
			a++
		}
	}
	fmt.Println(a)
}

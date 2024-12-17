package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func run(a, b, c int, pgm []int) {
	res := []int{}
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
			res = append(res, curr%8)
		case 6:
			b = a >> curr
		case 7:
			c = a >> curr
		}
	}

	resStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(res)), ","), "[]")
	fmt.Println(resStr)
}

func main() {
	input, _ := os.ReadFile("../../input/day17.txt")

	m := regexp.MustCompile(`[\d,]+`).FindAllString(string(input), -1)

	a, _ := strconv.Atoi(m[0])
	b, _ := strconv.Atoi(m[1])
	c, _ := strconv.Atoi(m[2])
	var pgm []int
	json.Unmarshal([]byte("["+m[3]+"]"), &pgm)

	run(a, b, c, pgm)
}

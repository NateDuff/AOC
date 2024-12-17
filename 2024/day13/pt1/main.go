package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Coordinates struct {
	X, Y int
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// solveLinearEquation accepts the coeffecients of two linear equations:
//
//	a1x + b1y = c1
//	a2x + b2y = c2
//
// the function retruns two values:
//
// boolen - to indicate if the value of X and Y are whole numbers or not.
// coordinates - actual values for X and Y
func SolveLinearEquation(a, b, c Coordinates) (bool, Coordinates) {
	x := ((b.X * (-c.Y)) - (b.Y * (-c.X))) / ((a.X * b.Y) - (a.Y * b.X))
	y := (((-c.X) * a.Y) - ((-c.Y) * a.X)) / ((a.X * b.Y) - (a.Y * b.X))
	if (((b.X*(-c.Y))-(b.Y*(-c.X)))%((a.X*b.Y)-(a.Y*b.X)) == 0) &&
		((((-c.X)*a.Y)-((-c.Y)*a.X))%((a.X*b.Y)-(a.Y*b.X)) == 0) {
		return true, Coordinates{x, y}
	}
	return false, Coordinates{x, y}
}

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

func fetchSliceOfIntsInString(line string) []int {
	nums := []int{}
	var build strings.Builder
	isNegative := false
	for _, char := range line {
		if unicode.IsDigit(char) {
			build.WriteRune(char)
		}

		if char == '-' {
			isNegative = true
		}

		if (char == ' ' || char == ',' || char == '~' || char == '|') && build.Len() != 0 {
			localNum, err := strconv.ParseInt(build.String(), 10, 64)
			if err != nil {
				panic(err)
			}
			if isNegative {
				localNum *= -1
			}
			nums = append(nums, int(localNum))
			build.Reset()
			isNegative = false
		}
	}
	if build.Len() != 0 {
		localNum, err := strconv.ParseInt(build.String(), 10, 64)
		if err != nil {
			panic(err)
		}
		if isNegative {
			localNum *= -1
		}
		nums = append(nums, int(localNum))
		build.Reset()
	}
	return nums
}

func getPrizes(input []string) [][]Coordinates {
	values := []Coordinates{}
	prizes := [][]Coordinates{}
	for _, line := range input {
		if len(line) == 0 {
			continue
		}
		nums := fetchSliceOfIntsInString(line)
		values = append(values, Coordinates{nums[0], nums[1]})

		if strings.Contains(line, "Prize") {
			prizes = append(prizes, values)
			values = []Coordinates{}
		}
	}
	return prizes
}

// getButtonPressesIfValid checks if there is a valid number
// of button presses to achieve the target.
//
// Consider A and B to be the final number of buttom presses for
// buttons A and B respectively.
//
// 1. Its only possible to acheive the target if A and B are integers.
// (We can't press a button in fraction).
//
// 2. To find values of A and B, we can get it by solving the linear
// equations (explained why below).
func getButtonPressesIfValid(a, b, final Coordinates, add int) (bool, Coordinates) {
	// Consider this:
	// On pressing button A: X moves by A1 and Y by A2
	// On pressing button B: X moves by B1 and Y by B2
	// The final location needed is (C1, C2).
	// If X and Y be the button presses for A and B respectively, then:
	// A1*X + B1*Y = C1
	// A2*X + B2*Y = C2
	// the above are just linear equations!
	// we can get the values of Y and Y easily using Crammer's Rule.
	final.X += add
	final.Y += add
	return SolveLinearEquation(a, b, final)
}

func getTokenCount(input [][]Coordinates) (int, int) {
	count1, count2 := 0, 0
	for _, prize := range input {
		isValid1, tokens1 := getButtonPressesIfValid(prize[0], prize[1], prize[2], 0)
		if isValid1 {
			count1 += tokens1.X*3 + tokens1.Y
		}

		isValid2, tokens2 := getButtonPressesIfValid(prize[0], prize[1], prize[2], 10000000000000)
		if isValid2 {
			count2 += tokens2.X*3 + tokens2.Y

		}
	}
	return count1, count2
}

func main() {
	input, err := ReadLines("../../input/day13.txt")
	if err != nil {
		panic(err)
	}
	ans1, ans2 := getTokenCount(getPrizes(input))
	fmt.Println(ans1)
	fmt.Println(ans2)
}

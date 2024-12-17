package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

// Robot struct represents a robot with position P and velocity V
type Robot struct {
	P, V image.Point
}

// sgn function returns the sign of an integer
func sgn(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}

func main() {
	// Read input from file
	input, _ := os.ReadFile("../../input/day14.txt")
	// Define the area as a rectangle
	area := image.Rectangle{image.Point{0, 0}, image.Point{101, 103}}

	// Initialize robots slice and quads map
	robots, quads := []Robot{}, map[image.Point]int{}
	// Parse input and initialize robots
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var r Robot
		// Scan the input line into robot's position and velocity
		fmt.Sscanf(s, "p=%d,%d v=%d,%d", &r.P.X, &r.P.Y, &r.V.X, &r.V.Y)
		robots = append(robots, r)
		// Update robot's position and calculate its quadrant
		r.P = r.P.Add(r.V.Mul(100)).Mod(area)
		quads[image.Point{sgn(r.P.X - area.Dx()/2), sgn(r.P.Y - area.Dy()/2)}]++
	}
	// Print the product of the number of robots in each quadrant
	fmt.Println(quads[image.Point{-1, -1}] * quads[image.Point{1, -1}] * quads[image.Point{1, 1}] * quads[image.Point{-1, 1}])

	// Simulate the movement of robots
	for t := 1; ; t++ {
		seen := map[image.Point]struct{}{}
		for i := range robots {
			// Update robot's position
			robots[i].P = robots[i].P.Add(robots[i].V).Mod(area)
			seen[robots[i].P] = struct{}{}
		}
		// Check if all robots are in unique positions
		if len(seen) == len(robots) {
			fmt.Println(t)
			break
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//p1()
	p2()
}

func p1() {
	mX, mY, towerLocs := readInput_8("../../input/day8.txt")
	antiNodes := findAntiNodes(mX, mY, towerLocs)
	fmt.Println(len(antiNodes))
}
func p2() {
	mX, mY, towerLocs := readInput_8("../../input/day8.txt")
	antiNodes := findAntiNodes_w_Harmonics(mX, mY, towerLocs)
	fmt.Println(len(antiNodes))
}

type locs struct {
	x int
	y int
}

func readInput_8(fileName string) (int, int, map[string][]locs) { //maxX, maxY, towerLocs map[frequency][]locs
	var mapRep [][]string

	file, _ := os.Open(fileName)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		temp := strings.Split(line, "")
		mapRep = append(mapRep, temp)
	}

	maxX := len(mapRep[0])
	maxY := len(mapRep)

	towerLocs := getTowerLocsFromMap(mapRep)

	return maxX, maxY, towerLocs
}

func getTowerLocsFromMap(mapIn [][]string) map[string][]locs {
	towerLocs := make(map[string][]locs)
	for i := 0; i < len(mapIn); i++ {
		for j := 0; j < len(mapIn[i]); j++ {
			if mapIn[i][j] != "." {
				towerLocs[mapIn[i][j]] = append(towerLocs[mapIn[i][j]], locs{x: j, y: i})
			}
		}
	}
	return towerLocs
}

func findAntiNodes(maxX int, maxY int, towerLocs map[string][]locs) map[locs]bool {
	antiNodeLocs := make(map[locs]bool)

	for _, values := range towerLocs {
		for i := 0; i < len(values)-1; i++ {
			for j := i + 1; j < len(values); j++ {
				distX := values[j].x - values[i].x
				distY := values[j].y - values[i].y

				antiNode_1 := locs{x: values[i].x - distX, y: values[i].y - distY}
				antiNode_2 := locs{x: values[j].x + distX, y: values[j].y + distY}

				if antiNode_1.x < maxX && antiNode_1.x >= 0 && antiNode_1.y < maxY && antiNode_1.y >= 0 {
					antiNodeLocs[antiNode_1] = true
				}
				if antiNode_2.x < maxX && antiNode_2.x >= 0 && antiNode_2.y < maxY && antiNode_2.y >= 0 {
					antiNodeLocs[antiNode_2] = true
				}
			}
		}
	}
	return antiNodeLocs
}

func findAntiNodes_w_Harmonics(maxX int, maxY int, towerLocs map[string][]locs) map[locs]bool {
	antiNodeLocs := make(map[locs]bool)

	for _, values := range towerLocs {
		for i := 0; i < len(values)-1; i++ {
			for j := i + 1; j < len(values); j++ {
				distX := values[j].x - values[i].x
				distY := values[j].y - values[i].y

				p1Shifted := createShiftedLocs(values[i], distX, distY, maxX, maxY)
				p2Shifted := createShiftedLocs(values[i], distX*-1, distY*-1, maxX, maxY)

				for k := 0; k < len(p1Shifted); k++ {
					antiNodeLocs[p1Shifted[k]] = true
				}
				for k := 0; k < len(p2Shifted); k++ {
					antiNodeLocs[p2Shifted[k]] = true
				}
				if len(p1Shifted) == 1 {
					antiNodeLocs[values[j]] = true
				}
				if len(p2Shifted) == 1 {
					antiNodeLocs[values[i]] = true
				}
				if len(p1Shifted) >= 2 || len(p2Shifted) >= 2 {
					antiNodeLocs[values[i]] = true
					antiNodeLocs[values[j]] = true
				}
			}
		}
	}
	return antiNodeLocs
}

func createShiftedLocs(starting locs, distX int, distY int, maxX int, maxY int) []locs {
	lastLoc := starting
	lastLocValid := true
	var locsOut []locs

	for lastLocValid {
		lastLoc.x -= distX
		lastLoc.y -= distY

		if lastLoc.x < maxX && lastLoc.x >= 0 && lastLoc.y < maxY && lastLoc.y >= 0 {
			locsOut = append(locsOut, lastLoc)
		} else {
			lastLocValid = false
		}
	}

	return locsOut
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	r int
	c int
}

type Warehouse struct {
	moveSeq string
	boxes   map[Pair]struct{}
	robot   Pair
	walls   map[Pair]struct{}
	width   int
	height  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readData(fileName string) Warehouse {
	file, err := os.Open(fileName)
	check(err)
	scanner := bufio.NewScanner(file)
	var house Warehouse
	house.boxes = make(map[Pair]struct{})
	house.walls = make(map[Pair]struct{})
	scanner.Scan()
	topWalls := scanner.Text()
	for i := range len(topWalls) {
		house.walls[Pair{0, i}] = struct{}{}
	}
	r := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == topWalls {
			for i := range len(topWalls) {
				house.walls[Pair{r, i}] = struct{}{}
			}
			break
		}

		chars := []rune(line)
		for col, char := range chars {
		switchState:
			switch char {
			case '#':
				house.walls[Pair{r, col}] = struct{}{}
				break switchState
			case 'O':
				house.boxes[Pair{r, col}] = struct{}{}
				break switchState
			case '@':
				house.robot = Pair{r, col}
				break switchState
			default:
				break switchState
			}
		}
		r++
	}
	house.height = r + 1
	house.width = len(topWalls)

	scanner.Scan()
	for scanner.Scan() {
		house.moveSeq += scanner.Text()
	}

	return house
}

func canBoxMove(house *Warehouse, box Pair, dir rune) bool {
	nextCoords := getNextPair(box, dir)
	_, ok := house.walls[nextCoords]
	if ok {
		return false // wall
	}
	_, ok = house.boxes[nextCoords]
	if ok {
		return canBoxMove(house, nextCoords, dir)
	}
	return true // nothing in front
}

func moveBoxes(house *Warehouse, box Pair, dir rune) {
	nextCoords := getNextPair(box, dir)
	_, ok := house.boxes[nextCoords]
	if ok {
		moveBoxes(house, nextCoords, dir)
	}
	delete(house.boxes, box)
	house.boxes[nextCoords] = struct{}{}
}

func move(house *Warehouse, dir rune) {
	nextCoords := getNextPair(house.robot, dir)
	_, ok := house.walls[nextCoords]
	if ok {
		return // wall in front
	}
	_, ok = house.boxes[nextCoords]
	if ok {
		if canBoxMove(house, nextCoords, dir) { // boxes in front
			moveBoxes(house, nextCoords, dir)
		} else {
			return
		}
	}
	house.robot = nextCoords // nothing in front
}

func run(house Warehouse) uint {
	moves := []rune(house.moveSeq)
	for _, char := range moves {
		move(&house, char)
	}

	ret := uint(0)

	for key := range house.boxes {
		ret += 100*uint(key.r) + uint(key.c)
	}

	return ret
}

func getNextPair(p Pair, dir rune) Pair {
	nextPair := Pair{p.r, p.c}
	if dir == '^' {
		nextPair.r--
	} else if dir == '>' {
		nextPair.c++
	} else if dir == 'v' {
		nextPair.r++
	} else {
		nextPair.c--
	}
	return nextPair
}

func main() {
	house := readData("../../input/day15.txt")
	fmt.Println(run(house))
}

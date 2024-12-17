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

type BigBox struct {
	left  Pair
	right Pair
}

type BigWarehouse struct {
	moveSeq  string
	boxes    map[BigBox]struct{}
	boxParts map[Pair]BigBox
	robot    Pair
	walls    map[Pair]struct{}
	width    int
	height   int
}

func NewBigWarehouse(house Warehouse) BigWarehouse {
	var newHouse BigWarehouse

	newHouse.height = house.height
	newHouse.width = house.width * 2
	newHouse.moveSeq = house.moveSeq
	newHouse.boxParts = make(map[Pair]BigBox)
	newHouse.walls = make(map[Pair]struct{})
	newHouse.boxes = make(map[BigBox]struct{})

	newHouse.robot = Pair{house.robot.r, house.robot.c * 2}
	for wall := range house.walls {
		wall1 := Pair{wall.r, wall.c * 2}
		wall2 := Pair{wall1.r, wall1.c + 1}
		newHouse.walls[wall1] = struct{}{}
		newHouse.walls[wall2] = struct{}{}
	}

	for box := range house.boxes {
		left := Pair{box.r, box.c * 2}
		right := Pair{box.r, left.c + 1}
		bigBox := BigBox{left, right}
		newHouse.boxParts[left] = bigBox
		newHouse.boxParts[right] = bigBox
		newHouse.boxes[bigBox] = struct{}{}
	}

	return newHouse
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

func canBigBoxMove(house *BigWarehouse, side Pair, dir rune) bool {
	canMove := true
	bb := house.boxParts[side]
	left, right := bb.left, bb.right
	leftNext := getNextPair(left, dir)
	rightNext := getNextPair(right, dir)

	_, lOK := house.walls[leftNext]
	_, rOK := house.walls[rightNext]
	if lOK || rOK {
		return false
	}

	if dir == '<' {
		_, lOK := house.boxParts[leftNext]
		if lOK {
			canMove = canBigBoxMove(house, leftNext, dir)
		}
		return canMove
	}

	if dir == '>' {
		_, rOK := house.boxParts[rightNext]
		if rOK {
			canMove = canBigBoxMove(house, rightNext, dir)
		}
		return canMove
	}

	// only runs if moving up or down
	bbL, lOk := house.boxParts[leftNext]
	bbR, rOk := house.boxParts[rightNext]

	if lOk {
		canMove = canMove && canBigBoxMove(house, leftNext, dir)
	}
	if rOk && bbL != bbR {
		canMove = canMove && canBigBoxMove(house, rightNext, dir)
	}

	return canMove
}

func bigBoxMove(house *BigWarehouse, side Pair, dir rune) {
	bb := house.boxParts[side]
	left, right := bb.left, bb.right
	leftNext, rightNext := getNextPair(left, dir), getNextPair(right, dir)

	if dir == '<' {
		_, lOK := house.boxParts[leftNext]
		if lOK {
			bigBoxMove(house, leftNext, dir)
		}

		delete(house.boxes, bb)
		delete(house.boxParts, left)
		delete(house.boxParts, right)

		bb.right = bb.left
		bb.left = leftNext
		house.boxes[bb] = struct{}{}
		house.boxParts[left] = bb
		house.boxParts[leftNext] = bb
		return
	}

	if dir == '>' {
		_, rOK := house.boxParts[rightNext]
		if rOK {
			bigBoxMove(house, rightNext, dir)
		}
		delete(house.boxes, bb)
		delete(house.boxParts, left)
		delete(house.boxParts, right)

		bb.left = bb.right
		bb.right = rightNext

		house.boxes[bb] = struct{}{}
		house.boxParts[right] = bb
		house.boxParts[rightNext] = bb
		return
	}

	bbL, lOK := house.boxParts[leftNext]
	bbR, rOk := house.boxParts[rightNext]

	if lOK {
		bigBoxMove(house, leftNext, dir)
	}
	if rOk && bbL != bbR {
		bigBoxMove(house, rightNext, dir)
	}
	delete(house.boxes, bb)
	delete(house.boxParts, left)
	delete(house.boxParts, right)

	bb.left = leftNext
	bb.right = rightNext

	house.boxes[bb] = struct{}{}
	house.boxParts[leftNext] = bb
	house.boxParts[rightNext] = bb
}

func bigMove(house *BigWarehouse, dir rune) {
	nextCoords := getNextPair(house.robot, dir)

	_, ok := house.walls[nextCoords]
	if ok {
		return
	}

	_, ok = house.boxParts[nextCoords]
	if ok {
		if canBigBoxMove(house, nextCoords, dir) {
			bigBoxMove(house, nextCoords, dir)
		} else {
			return
		}
	}
	house.robot = nextCoords
}

func run(house BigWarehouse) uint {
	moves := []rune(house.moveSeq)
	for _, char := range moves {
		bigMove(&house, char)
	}
	sum := uint(0)
	for bb := range house.boxes {
		sum += 100*uint(bb.left.r) + uint(bb.left.c)
	}
	return sum
}

func main() {
	house := readData("../../input/day15.txt")
	wareHouse := NewBigWarehouse(house)
	fmt.Println(run(wareHouse))
}

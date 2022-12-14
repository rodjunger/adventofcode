package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	right direction = "R"
	left  direction = "L"
	up    direction = "U"
	down  direction = "D"
)

type position struct {
	x, y int
}

func (p *position) move(d direction) {
	switch d {
	case right:
		p.y++
	case left:
		p.y--
	case up:
		p.x--
	case down:
		p.x++
	}
}

func (p *position) distanceFrom(other *position) position {
	return position{x: p.x - other.x, y: p.y - other.y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func planMoves(head *position, tail *position) []direction {
	var (
		out  []direction
		dist = tail.distanceFrom(head)
	)
	//head is more than one block up or is one block up and two left or right
	if dist.x > 1 || dist.x == 1 && (abs(dist.y) == 2) {
		out = append(out, up)
	}
	//head is more than one block down or is one block down and left or right
	if dist.x < -1 || dist.x == -1 && (abs(dist.y) == 2) {
		out = append(out, down)
	}

	//head is more than one block to the left or one left and two up or down
	if dist.y > 1 || dist.y == 1 && (abs(dist.x) == 2) {
		out = append(out, left)
	}

	//head is more than one block to the right or one right and two up or down
	if dist.y < -1 || dist.y == -1 && (abs(dist.x) == 2) {
		out = append(out, right)
	}

	return out
}

func (p *position) hash() string {
	return string(p.x) + string(p.y)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	visitedMap := map[string]struct{}{}
	headPosition := &position{999999, 999999} // using big numbers because small numbers break my algorithm (coords become negative eventually)
	tailPosition := &position{999999, 999999} // using big numbers because small numbers break my algorithm (coords become negative eventually)
	visitedMap[tailPosition.hash()] = struct{}{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		fmt.Println(parts)
		moveDirection := direction(parts[0])
		ammount, _ := strconv.Atoi(parts[1])
		for i := 0; i < ammount; i++ {
			headPosition.move(moveDirection)
			fmt.Println("new head position", headPosition)
			tailMoves := planMoves(headPosition, tailPosition)
			fmt.Println("tail moves:", tailMoves)
			for _, move := range tailMoves {
				tailPosition.move(move)
				fmt.Println("new tail position", tailPosition)
			}
			visitedMap[tailPosition.hash()] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(visitedMap))

}

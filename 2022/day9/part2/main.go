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

	knots := [10]*position{}

	for i := range knots {
		knots[i] = &position{999999, 999999} // should have a good way to calculate the size of the map but this works
	}

	visitedMap[knots[len(knots)-1].hash()] = struct{}{}

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		fmt.Println(parts)
		moveDirection := direction(parts[0])
		ammount, _ := strconv.Atoi(parts[1])
		for i := 0; i < ammount; i++ {
			knots[0].move(moveDirection)
			fmt.Println("new head position", knots[0])
			for j := 1; j < len(knots); j++ {
				tailMoves := planMoves(knots[j-1], knots[j])
				fmt.Println("tail moves:", tailMoves)
				for _, move := range tailMoves {
					knots[j].move(move)
				}
			}
			visitedMap[knots[len(knots)-1].hash()] = struct{}{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(visitedMap))

}

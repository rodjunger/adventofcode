package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type blockType int

const (
	air blockType = iota
	rock
	sand
)

// default return value is air, so only need to set rock and sand
type caveMap struct {
	cave     map[pos]blockType
	maxDepth int
}

// x is column, y is row
type pos struct {
	x, y int
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (p pos) generateFallingDirections() []pos {
	down, horizontalLeft, horizontalRight := pos{p.x, p.y + 1}, pos{p.x - 1, p.y + 1}, pos{p.x + 1, p.y + 1}
	return []pos{down, horizontalLeft, horizontalRight}
}

// returns true on success, false if it fell forever
func (c *caveMap) simulateSand() bool {
	sandLocation := pos{500, 0}

	for sandLocation.y <= c.maxDepth {
		var (
			possibleLocations = sandLocation.generateFallingDirections()
			foundLocation     bool
		)
		for i := 0; i < len(possibleLocations) && !foundLocation; i++ {
			loc := possibleLocations[i]
			if c.cave[loc] == air {
				sandLocation = loc
				foundLocation = true
			}
		}
		//Nowhere to go
		if !foundLocation {
			c.cave[sandLocation] = sand
			return true
		}
	}
	return false
}

func (c *caveMap) writeVertical(x, startY, endY int) {
	currentY, endY := Min(startY, endY), Max(startY, endY)
	for ; currentY <= endY; currentY++ {
		c.cave[pos{x, currentY}] = rock
	}
}

func (c *caveMap) writeHorizontal(y, startX, endX int) {
	currentX, endX := Min(startX, endX), Max(startX, endX)
	for ; currentX <= endX; currentX++ {
		c.cave[pos{currentX, y}] = rock
	}
}

func (c *caveMap) writeLine(start, end pos) {
	maxY := Max(start.y, end.y)
	if maxY > c.maxDepth {
		c.maxDepth = maxY
	}

	if start.x == end.x {
		c.writeVertical(start.x, start.y, end.y)
	} else {
		c.writeHorizontal(start.y, start.x, end.x)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cave := caveMap{
		cave: make(map[pos]blockType),
	}

	for scanner.Scan() {
		line := scanner.Text()
		points := []pos{}

		for _, coord := range strings.Split(line, " -> ") {
			parts := strings.Split(coord, ",")
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			points = append(points, pos{x, y})
		}

		for i := 0; i < len(points)-1; i++ {
			cave.writeLine(points[i], points[i+1])
		}
	}

	var sandsFallen int

	for cave.simulateSand() {
		sandsFallen++
	}

	fmt.Println(sandsFallen)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

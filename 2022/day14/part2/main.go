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
	cave     map[[2]int]blockType
	maxDepth int
}

// x is column, y is row
type pos struct {
	x, y int
}

func (p pos) toStruct() [2]int {
	return [2]int{p.x, p.y}
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

func (p pos) generateFallingDirections() [3][2]int {
	down, horizontalLeft, horizontalRight := pos{p.x, p.y + 1}.toStruct(), pos{p.x - 1, p.y + 1}.toStruct(), pos{p.x + 1, p.y + 1}.toStruct()
	return [3][2]int{down, horizontalLeft, horizontalRight}
}

// returns true on success, false if it fell forever
func (c *caveMap) simulateSand() bool {
	sandLocation := pos{500, 0}

	if c.cave[sandLocation.toStruct()] == sand {
		return false
	}

	for sandLocation.y < c.maxDepth+2 {
		var (
			possibleLocations = sandLocation.generateFallingDirections()
			foundLocation     bool
		)
		for i := 0; i < len(possibleLocations) && !foundLocation; i++ {
			loc := possibleLocations[i]
			var block blockType
			//Simulate the rock bottom
			if loc[1] == c.maxDepth+2 {
				block = rock
			} else {
				block = c.cave[loc]
			}
			if block == air {
				sandLocation = pos{loc[0], loc[1]}
				foundLocation = true
			}
		}
		//Nowhere to go
		if !foundLocation {
			c.cave[sandLocation.toStruct()] = sand
			return true
		}
	}
	//Reached max depth, shouldn't happen on this version
	return false
}

func (c *caveMap) writeVertical(x, startY, endY int) {
	currentY, endY := Min(startY, endY), Max(startY, endY)
	if endY > c.maxDepth {
		c.maxDepth = endY
	}
	for ; currentY <= endY; currentY++ {
		//fmt.Println("writing at", x, currentY)
		c.cave[[2]int{x, currentY}] = rock
	}
}

func (c *caveMap) writeHorizontal(y, startX, endX int) {
	currentX, endX := Min(startX, endX), Max(startX, endX)
	for ; currentX <= endX; currentX++ {
		//fmt.Println("writing at", currentX, y)
		c.cave[[2]int{currentX, y}] = rock
	}
}

func (c *caveMap) writeLine(start, end pos) {
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
		cave: make(map[[2]int]blockType),
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sandsFallen int

	for cave.simulateSand() {
		sandsFallen++
	}

	fmt.Println(sandsFallen)

}

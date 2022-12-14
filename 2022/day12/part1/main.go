package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	row    int
	column int
	father *point
}

func validDirection(current, desired point, heightMap [][]rune, visited map[int]map[int]bool) bool {
	switch {
	case desired.column > len(heightMap[0])-1,
		desired.column < 0,
		desired.row > len(heightMap)-1,
		desired.row < 0,
		heightMap[desired.row][desired.column] > heightMap[current.row][current.column]+1,
		visited[desired.row][desired.column]:
		return false
	default:
		return true
	}
}

func findShortestPath(start, end point, heightMap [][]rune, nSteps int) int {
	visited := map[int]map[int]bool{}

	for i := range heightMap {
		visited[i] = make(map[int]bool)
	}

	var (
		//use a large buffered channel as queue, the size can be calculated based on the number of nodes but this works lol
		queue       = make(chan point, 1000)
		currentEdge point
	)

	visited[start.row][start.column] = true

	queue <- start

	for {
		currentEdge = <-queue
		if currentEdge.row == end.row && currentEdge.column == end.column {
			var steps int
			for ; currentEdge.father != nil; steps++ {
				currentEdge = *currentEdge.father
			}
			return steps
		}

		edges := []point{
			{row: currentEdge.row, column: currentEdge.column + 1},
			{row: currentEdge.row, column: currentEdge.column - 1},
			{row: currentEdge.row - 1, column: currentEdge.column},
			{row: currentEdge.row + 1, column: currentEdge.column},
		}

		for _, edge := range edges {
			valid := validDirection(currentEdge, edge, heightMap, visited)
			if valid {
				//make copies to have unique objects/pointers
				cur := currentEdge
				new := edge
				new.father = &cur
				visited[new.row][new.column] = true
				queue <- new
			}
		}
		if len(queue) == 0 {
			return 0
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		heightMap  [][]rune
		start, end point
	)

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		heightMap = append(heightMap, []rune(line))
		for i, r := range line {
			//Converting to rune and back to string is kind meh
			if string(r) == "S" {
				heightMap[row][i] = 'a'
				start = point{row: row, column: i}
			} else if string(r) == "E" {
				heightMap[row][i] = 'z'
				end = point{row: row, column: i}
			}
		}
	}

	fmt.Println(findShortestPath(start, end, heightMap, 0))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

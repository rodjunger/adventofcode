package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type tree struct {
	height int
}

func scenicScore(treeMap [][]*tree, row, column int) int {
	var (
		maxRow                                    = len(treeMap)
		maxColumn                                 = len(treeMap[0])
		height                                    = treeMap[row][column].height
		rightScore, leftScore, upScore, downScore int
	)
	//Check to the right
	for i := column + 1; i < maxColumn; i++ {
		rightScore++
		if treeMap[row][i].height >= height {
			break
		}
	}
	//To the left
	for i := column - 1; i >= 0; i-- {
		leftScore++
		if treeMap[row][i].height >= height {
			break

		}
	}
	//Going down
	for i := row + 1; i < maxRow; i++ {
		downScore++
		if treeMap[i][column].height >= height {
			break
		}
	}
	//Going up
	for i := row - 1; i >= 0; i-- {
		upScore++
		if treeMap[i][column].height >= height {
			break
		}
	}

	return leftScore * rightScore * upScore * downScore
}

func walkMap(treeMap [][]*tree) int {
	var out int
	for i, row := range treeMap {
		for j := range row {
			score := scenicScore(treeMap, i, j)
			if score > out {
				out = score
			}
			//fmt.Println("score", i, j, "=", score)
		}
	}
	return out
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var treeMap [][]*tree

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		treeMap = append(treeMap, []*tree{})
		for _, letter := range line {
			height, _ := strconv.Atoi(string(letter))
			treeMap[i] = append(treeMap[i], &tree{height: height})
		}
		//fmt.Println(treeMap)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := walkMap(treeMap)
	fmt.Println(result)

}

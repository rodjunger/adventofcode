package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type tree struct {
	height  int
	visible bool
}

type direction int

const (
	leftToRight direction = iota
	rightToLeft
	topToBottom
	bottomToTop
)

func walkHorizontal(treeLine []*tree, direction direction) {
	switch direction {
	case rightToLeft:
		highest := treeLine[len(treeLine)-1].height - 1
		for i := len(treeLine) - 1; i != 0; i-- {
			if treeLine[i].height > highest {
				treeLine[i].visible = true
				highest = treeLine[i].height
			}
		}
	case leftToRight:
		highest := treeLine[0].height - 1
		for i := range treeLine {
			if treeLine[i].height > highest {
				treeLine[i].visible = true
				highest = treeLine[i].height
			}
		}
	}
}

func walkVertical(treeMap [][]*tree, column int, direction direction) {
	switch direction {
	case bottomToTop:
		highest := treeMap[len(treeMap)-1][column].height - 1
		for i := len(treeMap) - 1; i != 0; i-- {
			height := treeMap[i][column].height
			if height > highest {
				treeMap[i][column].visible = true
				highest = height
			}
		}
	case topToBottom:
		highest := treeMap[0][column].height - 1
		for i := range treeMap {
			height := treeMap[i][column].height
			if height > highest {
				treeMap[i][column].visible = true
				highest = height
			}
		}
	}
}

func walkMap(treeMap [][]*tree) int {
	//Could use goroutines to do every line and every direction in parallel, one after another
	for i := range treeMap {
		walkHorizontal(treeMap[i], leftToRight)
		walkHorizontal(treeMap[i], rightToLeft)
	}

	for i := range treeMap[0] {
		walkVertical(treeMap, i, topToBottom)
		walkVertical(treeMap, i, bottomToTop)
	}

	var out int
	for _, row := range treeMap {
		for _, tree := range row {
			if tree.visible {
				out++
			}
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

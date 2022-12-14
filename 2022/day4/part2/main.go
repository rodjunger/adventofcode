package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type section struct {
	start  int
	finish int
}

func (s section) simplyOverlaps(other section) bool {
	switch {
	case s.start >= other.start && s.start <= other.finish,
		s.finish >= other.start && s.finish <= other.start:
		return true
	default:
		return false
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numberOfOverlaps int

	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")
		firstElve, secondElve := strings.Split(elves[0], "-"), strings.Split(elves[1], "-")
		firstElveStart, _ := strconv.Atoi(firstElve[0])
		firstElveFinish, _ := strconv.Atoi(firstElve[1])
		secondElveStart, _ := strconv.Atoi(secondElve[0])
		secondElveFinish, _ := strconv.Atoi(secondElve[1])

		e1 := section{firstElveStart, firstElveFinish}
		e2 := section{secondElveStart, secondElveFinish}
		if e1.simplyOverlaps(e2) || e2.simplyOverlaps(e1) {
			numberOfOverlaps += 1
		}
	}

	fmt.Println(numberOfOverlaps)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type RPSOptions int

const (
	invalid RPSOptions = -1
	rock    RPSOptions = iota
	paper
	scisors
)

type RPSResult int

const (
	win RPSResult = iota
	lose
	draw
)

// No erro checking because there should be no invalid input on AOC
func fromLetter(letter string) RPSOptions {
	switch letter {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scisors
	default:
		return invalid
	}
}

// Returns true if me beats oponent
func emulateRockPaperScisors(oponent, me RPSOptions) RPSResult {
	switch {
	case oponent == me:
		return draw
	case me == rock && oponent == scisors,
		me == scisors && oponent == paper,
		me == paper && oponent == rock:
		return win
	default:
		return lose
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var points int

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		oponent, me := fromLetter(parts[0]), fromLetter(parts[1])
		matchResult := emulateRockPaperScisors(oponent, me)
		switch matchResult {
		case win:
			points += 6
		case draw:
			points += 3
		}
		points += int(me)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(points)
}

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

func resultFromEncryptedString(letter string) RPSResult {
	switch letter {
	case "X":
		return lose
	case "Y":
		return draw
	case "Z":
		return win
	}
	panic("invalid encrypted result")
}

// Returns what hand would be needed to have the desired result against oponent
func optionForResult(oponent RPSOptions, result RPSResult) (me RPSOptions) {
	switch result {
	case draw:
		return oponent
	case win:
		switch oponent {
		case scisors:
			return rock
		case paper:
			return scisors
		case rock:
			return paper
		}
	case lose:
		switch oponent {
		case rock:
			return scisors
		case scisors:
			return paper
		case paper:
			return rock
		}
	}
	return invalid
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
		oponent := fromLetter(parts[0])
		desiredResult := resultFromEncryptedString(parts[1])
		myHand := optionForResult(oponent, desiredResult)
		switch desiredResult {
		case win:
			points += 6
		case draw:
			points += 3
		}
		points += int(myHand)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(points)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var sum int
	for scanner.Scan() {
		ruckSack := scanner.Text()
		firstCompartment := ruckSack[:len(ruckSack)/2]
		secondCompartment := ruckSack[len(ruckSack)/2:]

		hashMap := make(map[rune]bool, len(firstCompartment))

		for _, letter := range firstCompartment {
			hashMap[letter] = true
		}

		var repeats rune

	findLoop:
		for _, letter := range secondCompartment {
			if hashMap[letter] {
				repeats = letter
				break findLoop
			}
		}

		if unicode.IsLower(repeats) {
			sum += int(repeats) - 96
			fmt.Println(string(repeats), int(repeats)-96)
		} else {
			sum += int(repeats) - 38
			fmt.Println(string(repeats), int(repeats)-38)
		}

	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// Change this to return an error to make it more readable, not changed on part2
func letterDistanceFrom(val rune, in string) (int, error) {
	var (
		lastIndex int
		found     bool
	)
	for i, rn := range in {
		if val == rn {
			lastIndex = i
			found = true
		}
	}

	if !found {
		return 0, errors.New("val not in input string")
	}

	return len(in) - lastIndex, nil
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		sinal := scanner.Text()
		var count int
		for i, letter := range sinal {
			dist, err := letterDistanceFrom(letter, sinal[i-count:i])
			fmt.Println("Distance from", string(letter), "to", sinal[i-count:i], ":", dist, "count:", count)
			if err != nil {
				count += 1
			} else {
				count = dist
			}
			if count == 4 {
				fmt.Println(i + 1)
				return
			}
		}
	}
}

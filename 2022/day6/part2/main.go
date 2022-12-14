package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Distance between val and the last val found in the string in
func letterDistanceFrom(val rune, in string) int {
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
		return 0
	}

	return len(in) - lastIndex
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
			a := letterDistanceFrom(letter, sinal[i-count:i])
			fmt.Println("Distance from", string(letter), "to", sinal[i-count:i], ":", a, "count:", count)
			if a == 0 {
				count += 1
			} else {
				count = a
			}
			if count == 14 {
				fmt.Println(i + 1)
				return
			}
		}
	}
}

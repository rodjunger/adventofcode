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
	var lettersInFirstRuckSack, lettersInFirstAndSecondRuckSack map[rune]bool

	for i := 0; scanner.Scan(); i += 1 {
		ruckSack := scanner.Text()
		switch i % 3 {
		case 0:
			lettersInFirstRuckSack = make(map[rune]bool, len(ruckSack))

			for _, letter := range ruckSack {
				lettersInFirstRuckSack[letter] = true
			}
		case 1:
			lettersInFirstAndSecondRuckSack = make(map[rune]bool)
			for _, letter := range ruckSack {
				if lettersInFirstRuckSack[letter] {
					lettersInFirstAndSecondRuckSack[letter] = true
				}
			}
		case 2:
		thisFor:
			for _, letter := range ruckSack {
				if lettersInFirstAndSecondRuckSack[letter] {
					if unicode.IsLower(letter) {
						sum += int(letter) - 96
						fmt.Println(string(letter), int(letter)-96)
					} else {
						sum += int(letter) - 38
						fmt.Println(string(letter), int(letter)-38)
					}
					break thisFor
				}
			}
		}
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

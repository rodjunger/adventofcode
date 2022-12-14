package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Sorted from biggest to lowest for easier iteration
type sortedTop []int

// never increases the size of the slice
func (s sortedTop) insertSorted(newValue int) {
	for i, currentValue := range s {
		if newValue > currentValue {
			copy(s[i+1:], s[i:])
			s[i] = newValue
			return
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

	var current int

	top3 := make(sortedTop, 3)

	for scanner.Scan() {
		converted, err := strconv.Atoi(scanner.Text())
		if err != nil { // got the empty line
			top3.insertSorted(current)
			current = 0
		} else {
			current += converted
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var total int
	for _, val := range top3 {
		total += val
	}

	fmt.Println(total)
}

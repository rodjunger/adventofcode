package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var current, biggest int

	for scanner.Scan() {
		converted, err := strconv.Atoi(scanner.Text())
		if err != nil {
			if current > biggest {
				biggest = current
			}
			current = 0
		}
		current += converted
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(biggest)
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack []string

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) pop() (string, error) {
	if s.isEmpty() {
		return "", errors.New("pop from empty stack")
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value, nil
}

func (s *stack) push(value string) {
	*s = append(*s, value)
}

func createStacks(lines stack) []stack {
	var (
		shouldContinue     = true
		numberOfContainers int
		containers         []stack
	)
	for i := 0; shouldContinue; i++ {
		line, err := lines.pop()
		if err != nil {
			shouldContinue = false
			continue
		}
		switch i {
		case 0:
			numberOfContainers = (len(line) / 4) + 1
			containers = make([]stack, numberOfContainers)
		default:
			for j := 0; j < numberOfContainers; j++ {
				currentLetter := string(line[(j*4)+1])
				if currentLetter != " " {
					containers[j].push(currentLetter)
				}
			}
		}
	}
	return containers
}

func parseCommand(command string) (quantity, from, to int) {
	parts := strings.Split(command, " ")
	quantity, _ = strconv.Atoi(parts[1])
	from, _ = strconv.Atoi(parts[3])
	to, _ = strconv.Atoi(parts[5])
	return
}

func handleMove(stacks []stack, command string) {
	quantity, from, to := parseCommand(command)
	var tempStack stack
	for i := 0; i < quantity; i++ {
		toMove, _ := stacks[from-1].pop() // we just know it will never fail because it's a challange with valid input
		tempStack.push(toMove)
	}
	for {
		elem, err := tempStack.pop()
		if err != nil {
			break
		}
		stacks[to-1].push(elem)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := "crates"
	var stacks []stack
	var lines stack

	for scanner.Scan() {
		line := scanner.Text()
		switch state {
		case "crates":
			if line == "" {
				stacks = createStacks(lines)
				state = "move"
			} else {
				lines.push(line)
			}
		case "move":
			handleMove(stacks, line)
		}
	}

	var final stack
	for _, st := range stacks {
		val, _ := st.pop()
		final.push(val)
	}

	fmt.Println(strings.Join(final, ""))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

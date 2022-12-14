package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	cycles            int
	X                 int
	signalStrengthSum int
}

func (c *cpu) executeInstruction(instruction string) {
	parts := strings.Split(instruction, " ")
	switch parts[0] {
	case "noop":
		c.incCycle()
	case "addx":
		register, _ := strconv.Atoi(parts[1])
		c.incCycle()
		c.incCycle()
		c.X += register
	}
}

func (c *cpu) incCycle() {
	c.cycles++
	if c.cycles%40 == 20 {
		c.signalStrengthSum += c.cycles * c.X
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cpu := &cpu{X: 1}
	for scanner.Scan() {
		instruction := scanner.Text()
		cpu.executeInstruction(instruction)
	}

	fmt.Println(cpu.signalStrengthSum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

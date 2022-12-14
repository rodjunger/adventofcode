package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type throw struct {
	worryLevel int
	to         int //monkey which item was throw to
}
type monkey struct {
	items          []int
	operation      operation
	testValue      int
	throwIdIfTrue  int
	throwIdIfFalse int
	inspectedItems int
}

func calculate(x, y int, operator string) int {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	}
	return 0 // would retorn err in normal code
}

func (m *monkey) inspectItems() []throw {
	var thrownItems []throw
	for i := range m.items {
		m.inspectedItems++
		worryLevel := m.items[i]
		var firstParam, secondParam int

		switch m.operation.firstParam {
		case "old":
			firstParam = worryLevel
		default:
			firstParam, _ = strconv.Atoi(m.operation.firstParam)
		}

		switch m.operation.secondParam {
		case "old":
			secondParam = worryLevel
		default:
			secondParam, _ = strconv.Atoi(m.operation.secondParam)
		}

		newWorryLevel := calculate(firstParam, secondParam, m.operation.operator)
		newWorryLevel /= 3

		to := m.throwIdIfFalse
		if newWorryLevel%m.testValue == 0 {
			to = m.throwIdIfTrue
		}

		thrownItems = append(thrownItems, throw{worryLevel: newWorryLevel, to: to})
	}
	//reset list because we just threw everything away
	m.items = []int{}
	return thrownItems
}

type operation struct {
	firstParam  string
	secondParam string
	operator    string
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var (
		monkeys   []*monkey
		curMonkey *monkey
	)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		switch i % 7 {
		case 0:
			curMonkey = &monkey{}
			monkeys = append(monkeys, curMonkey)
		case 1:
			worryLevels := strings.Split(parts[1], ", ")
			for _, level := range worryLevels {
				curLevel, _ := strconv.Atoi(level)
				curMonkey.items = append(curMonkey.items, curLevel)
			}
		case 2:
			operationParts := strings.Split(parts[1], " ")
			operationParts = operationParts[2:]
			curMonkey.operation = operation{
				firstParam:  operationParts[0],
				secondParam: operationParts[2],
				operator:    operationParts[1],
			}
		case 3:
			testParts := strings.Split(parts[1], " ")
			testValue, _ := strconv.Atoi(testParts[len(testParts)-1])
			curMonkey.testValue = testValue
		case 4:
			parts := strings.Split(parts[1], " ")
			monkeyId, _ := strconv.Atoi(parts[len(parts)-1])
			curMonkey.throwIdIfTrue = monkeyId
		case 5:
			parts := strings.Split(parts[1], " ")
			monkeyId, _ := strconv.Atoi(parts[len(parts)-1])
			curMonkey.throwIdIfFalse = monkeyId
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		for _, monkey := range monkeys {
			thrownItems := monkey.inspectItems()
			//fmt.Println("thrown items: ", thrownItems)
			for _, item := range thrownItems {
				monkeys[item.to].items = append(monkeys[item.to].items, item.worryLevel)
			}
		}
	}

	var totalInspects []int
	for i, monkey := range monkeys {
		fmt.Println("monkey", i, "inspected times", monkey.inspectedItems)
		totalInspects = append(totalInspects, monkey.inspectedItems)
	}

	sort.Ints(totalInspects)

	fmt.Println(totalInspects[len(totalInspects)-1] * totalInspects[len(totalInspects)-2])

}

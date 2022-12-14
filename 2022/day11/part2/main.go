package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type throw struct {
	worryLevel *big.Int
	to         int //monkey which item was throw to
}
type monkey struct {
	items          []*big.Int
	operation      operation
	testValue      *big.Int
	throwIdIfTrue  int
	throwIdIfFalse int
	inspectedItems int
	divider        *int
}

func calculate(x, y *big.Int, operator string) *big.Int {
	new := big.NewInt(0)
	switch operator {
	case "+":
		return new.Add(x, y)
	case "-":
		return new.Sub(x, y)
	case "*":
		return new.Mul(x, y)
	case "/":
		return new.Div(x, y)
	}
	return big.NewInt(0) // would retorn err in normal code
}

func (m *monkey) inspectItems() []throw {
	var thrownItems []throw
	for i := range m.items {
		m.inspectedItems++
		worryLevel := m.items[i]
		var firstParam, secondParam *big.Int

		switch m.operation.firstParam {
		case "old":
			firstParam = worryLevel
		default:
			res, _ := strconv.Atoi(m.operation.firstParam)
			firstParam = big.NewInt(int64(res))
		}

		switch m.operation.secondParam {
		case "old":
			secondParam = worryLevel
		default:
			res, _ := strconv.Atoi(m.operation.secondParam)
			secondParam = big.NewInt(int64(res))
		}

		//fmt.Println(firstParam, secondParam, m.operation.operator)
		newWorryLevel := calculate(firstParam, secondParam, m.operation.operator)
		//newWorryLevel /= 3

		//fmt.Println(newWorryLevel)

		to := m.throwIdIfFalse
		if big.NewInt(0).Mod(newWorryLevel, m.testValue).Int64() == 0 {
			to = m.throwIdIfTrue
		}

		newWorryLevel.Mod(newWorryLevel, big.NewInt(int64(*m.divider)))

		thrownItems = append(thrownItems, throw{worryLevel: newWorryLevel, to: to})
	}
	//reset list because we just threw everything away
	m.items = []*big.Int{}
	return thrownItems
}

type operation struct {
	firstParam  string
	secondParam string
	operator    string
}

// Had to google this one after wasting a lot of time implementing big.Int and noticing that it becomes exponentially slower with time
// Kept the big.Int code just for historical reasons
// Honestly, this is kind of a strange challange, it's weird that they expected me to know this theorem
// the change to the code is super small but requires very specific knowledge
// https://en.wikipedia.org/wiki/Chinese_remainder_theorem

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

	divider := 1

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		switch i % 7 {
		case 0:
			curMonkey = &monkey{divider: &divider}
			monkeys = append(monkeys, curMonkey)
		case 1:
			worryLevels := strings.Split(parts[1], ", ")
			for _, level := range worryLevels {
				curLevel, _ := strconv.Atoi(level)
				curMonkey.items = append(curMonkey.items, big.NewInt(int64(curLevel)))
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
			divider *= testValue
			curMonkey.testValue = big.NewInt(int64(testValue))
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

	for i := 0; i < 10000; i++ {
		if i%100 == 0 {
			fmt.Println("stage:", i)
		}
		for _, monkey := range monkeys {
			thrownItems := monkey.inspectItems()
			for _, item := range thrownItems {
				//fmt.Print(item.worryLevel, " ", item.to, " | ")
				monkeys[item.to].items = append(monkeys[item.to].items, big.NewInt(0).Set(item.worryLevel))
			}
			//fmt.Print("\n")
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

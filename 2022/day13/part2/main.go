package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"unicode"
)

type comparisonResult int

const (
	undecided comparisonResult = iota
	correct
	incorrect
)

// Using an interface was a horrible idea, don't write code sleepy
func compare(left, right []interface{}) comparisonResult {
	var leftType, rightType string
	//fmt.Print(left, "\n", right, "\n\n")

	switch {
	case len(left) == 0 && len(right) == 0:
		return undecided
	case len(left) == 0 && len(right) != 0:
		return correct
	case len(left) != 0 && len(right) == 0:
		return incorrect
	}

	switch left[0].(type) {
	case int:
		leftType = "int"
	case []interface{}:
		leftType = "list"
	}

	switch right[0].(type) {
	case int:
		rightType = "int"
	case []interface{}:
		rightType = "list"
	}

	var result comparisonResult
	switch {
	case leftType == rightType && leftType == "int":
		switch {
		case left[0].(int) < right[0].(int):
			return correct
		case left[0].(int) > right[0].(int):
			return incorrect
		default:
			return compare(left[1:], right[1:])
		}
		//Left for the memory
	case leftType == rightType && leftType == "list":
		result = compare(left[0].([]interface{}), right[0].([]interface{}))
	case leftType == "list" && rightType == "int":
		newRight := []interface{}{right[0].(int)}
		result = compare(left[0].([]interface{}), newRight)
	case leftType == "int" && rightType == "list":
		newLeft := []interface{}{left[0].(int)}
		result = compare(newLeft, right[0].([]interface{}))
	}

	if result != undecided {
		return result
	} else {
		return compare(left[1:], right[1:])
	}
}

func buildList(in []rune, index int) ([]interface{}, int) {

	thisSlice := []interface{}{}

	var digits []rune

	handleDigits := func() {
		if len(digits) > 0 {
			val, _ := strconv.Atoi(string(digits))
			thisSlice = append(thisSlice, val)
			digits = (digits)[:0]
		}
	}

	for index != len(in) {
		token := in[index]
		switch {
		case unicode.IsDigit(token):
			digits = append(digits, token)
			index++
		case token == '[':
			newSlice, nextIndex := buildList(in, index+1)
			thisSlice = append(thisSlice, newSlice)
			index = nextIndex
		case token == ']':
			handleDigits()
			return thisSlice, index + 1
		case token == ',':
			handleDigits()
			index++
		}
	}

	return thisSlice, index
}

type signalSorter struct {
	signals [][]interface{}
}

func (s *signalSorter) Len() int {
	return len(s.signals)
}

func (s *signalSorter) Less(i, j int) bool {
	return compare(s.signals[i], s.signals[j]) == correct
}

func (s *signalSorter) Swap(i, j int) {
	s.signals[i], s.signals[j] = s.signals[j], s.signals[i]
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lists [][]interface{}

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			res, _ := buildList([]rune(line), 1)
			lists = append(lists, res)
			//fmt.Println(res)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	dividerTwo, _ := buildList([]rune("[[2]]"), 1)
	dividerSix, _ := buildList([]rune("[[6]]"), 1)

	lists = append(lists, dividerTwo)
	lists = append(lists, dividerSix)

	sort.Sort(&signalSorter{lists})

	var twoIndex, sixIndex int
	for i, v := range lists {
		if compare(v, dividerTwo) == undecided {
			twoIndex = i + 1
		} else if compare(v, dividerSix) == undecided {
			sixIndex = i + 1
		}
	}
	fmt.Println(twoIndex * sixIndex)

}

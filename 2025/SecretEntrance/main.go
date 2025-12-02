package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type dial struct {
	CurrentValue int
	MaxValue     int
	CountZero    int
	PassedZero   int
}

func main() {
	Dial := dial{CurrentValue: 50, MaxValue: 100, CountZero: 0, PassedZero: 0}
	stopped, passed := Dial.FindCode("input.txt")
	fmt.Printf("Pt1. Stopped at 0: %d\n", stopped)
	fmt.Printf("Pt2. Passed 0: %d\n", passed)
}

func (Dial *dial) FindCode(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var firstChar byte
	var moveNum int
	var inputText string

	for scanner.Scan() {
		inputText = scanner.Text()
		firstChar = inputText[0]
		moveNum, _ = strconv.Atoi(inputText[1:])

		if firstChar == 'L' {
			Dial.spinLeft(moveNum)
		} else if firstChar == 'R' {
			Dial.spinRight(moveNum)
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error during scanning: %v", err)
		}
	}
	return Dial.CountZero, Dial.PassedZero
}

func (Dial *dial) spinRight(number int) int {
	newValue := (Dial.CurrentValue + number) % Dial.MaxValue

	passed := (Dial.CurrentValue + number) / Dial.MaxValue

	Dial.CurrentValue = newValue
	Dial.PassedZero += passed

	if Dial.CurrentValue == 0 {
		Dial.CountZero++
	}
	return Dial.CurrentValue
}

func (Dial *dial) spinLeft(number int) int {
	passed := 0

	if number >= Dial.CurrentValue {
		if Dial.CurrentValue == 0 {
			passed = number / Dial.MaxValue
		} else {
			passed = 1 + (number-Dial.CurrentValue)/Dial.MaxValue
		}
	}
	Dial.PassedZero += passed

	effectiveMove := number % Dial.MaxValue
	Dial.CurrentValue = (Dial.CurrentValue - effectiveMove + Dial.MaxValue) % Dial.MaxValue

	if Dial.CurrentValue == 0 {
		Dial.CountZero++
	}

	return Dial.CurrentValue
}

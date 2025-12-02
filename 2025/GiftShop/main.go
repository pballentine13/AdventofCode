package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	answerPT1 := fixGiftShop("input.txt")
	fmt.Println(answerPT1)
}

func fixGiftShop(filename string) int {
	idRanges := readInput(filename)

	sum := 0

	for _, idRange := range idRanges {
		start, end := parseRange(idRange)
		numbers := checkForInvalid(start, end)
		sum += sumArray(numbers)
	}

	return sum
}

func readInput(filename string) []string {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	inputString := string(input)
	itemRanges := strings.Split(inputString, ",")
	return itemRanges
}

func checkForInvalid(start int, end int) []int {
	invalid := []int{}
	for i := start; i <= end; i++ {
		s := len(strconv.Itoa(i))
		n := strconv.Itoa(i)

		if s%2 == 0 {
			firstHalf := n[:(s / 2)]
			secondHalf := n[s/2:]
			if firstHalf == secondHalf {
				invalid = append(invalid, i)
			}
		}
	}
	return invalid
}

func parseRange(idRange string) (int, int) {
	split := strings.Split(idRange, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])
	return start, end
}

func sumArray(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	day, _ := strconv.Atoi(args[0])
	part, _ := strconv.Atoi(args[1])
	mode := args[2]
	isTest := mode == "test"

	fileName := fmt.Sprintf("data/day%d_%s.txt", day, mode)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	switch day {
	case 1:
		output := dayOne(scanner, part)
		_run(part, isTest, [2]int{11, 31}, output)
	case 2:
		output := dayTwo(scanner, part)
		_run(part, isTest, [2]int{2, 4}, output)
	case 3:
		output := dayThree(scanner, part)
		_run(part, isTest, [2]int{161, 48}, output)
	case 4:
		output := dayFour(scanner, part)
		_run(part, isTest, [2]int{18, 9}, output)
	case 5:
		output := dayFive(scanner, part)
		_run(part, isTest, [2]int{143, 123}, output)
	}
}

func _run(part int, isTest bool, expectations [2]int, result int) {
	if isTest {
		var testExpectation int
		if part == 1 {
			testExpectation = expectations[0]
		} else {
			testExpectation = expectations[1]
		}
		if result != testExpectation {
			log.Fatalf("Invalid: expected %d, got %d", testExpectation, result)
		}
	} else {
		log.Printf("You got it tiger! Output: %d", result)
	}
}

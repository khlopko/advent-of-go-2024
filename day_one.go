package main

import (
	"bufio"
	"math"
	"slices"
	"strconv"
	"strings"
)

func dayOne(scanner *bufio.Scanner) int {
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")

		leftNum, _ := strconv.Atoi(nums[0])
		left = append(left, leftNum)

		rightNum, _ := strconv.Atoi(nums[1])
		right = append(right, rightNum)
	}

	if len(left) != len(right) {
		return 0
	}

	slices.Sort(left)
	slices.Sort(right)

	i := 0
	d := 0.0
	for i < len(left) {
		d += math.Abs(float64(left[i] - right[i]))
		i += 1
	}

	return int(d)
}

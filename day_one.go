package main

import (
	"bufio"
	"math"
	"slices"
	"strconv"
	"strings"
)

func dayOne(scanner *bufio.Scanner, part int) int {
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

	if part == 1 {
		i := 0
		d := 0.0
		for i < len(left) {
			d += math.Abs(float64(left[i] - right[i]))
			i += 1
		}

		return int(d)
	}

	if part == 2 {
		freq := make(map[int]int)
		k := 0
		for i := range left {
			lookingFor := left[i]
			for right[k] < lookingFor {
				k += 1
			}
			for right[k] == lookingFor {
				freq[lookingFor] += 1
				k += 1
			}
		}
		sum := 0
		for i := 0; i < len(left); i++ {
			sum += left[i] * freq[left[i]]
		}
		return sum
	}

	return -1
}

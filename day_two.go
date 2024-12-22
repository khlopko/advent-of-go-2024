package main

import (
	"bufio"
)

func dayTwo(scanner *bufio.Scanner, part int) int {
	result := 0
	for scanner.Scan() {
		row := scanner.Bytes()
		nums := readNums(row)

		if isSafelyIncreasing(append([]int{}, nums...)) || isSafelyDecreasing(append([]int{}, nums...)) {
			result += 1
		}
	}

	return result
}

func readNums(row []byte) []int {
	nums := []int{}
	curr := 0

	for _, c := range row {
		if c == 32 {
			nums = append(nums, curr)
			curr = 0
		} else {
			curr *= 10
			curr += int(c - 48)
		}
	}

	nums = append(nums, curr)

	return nums
}

func isSafelyIncreasing(nums []int) bool {
	tolerance := true
	i := 0
	k := 1
	for k < len(nums) {
		d := nums[k] - nums[i]
		if d < 1 || d > 3 {
			if tolerance == false {
				return false
			}
			nums = append(append([]int{}, nums[:i]...), nums[k:]...)
			tolerance = false
			continue
		}
		i += 1
		k += 1
	}

	return true
}

func isSafelyDecreasing(nums []int) bool {
	tolerance := true
	i := 0
	k := 1
	for k < len(nums) {
		d := nums[i] - nums[k]
		if d < 1 || d > 3 {
			if tolerance == false {
				return false
			}
			nums = append(append([]int{}, nums[:k]...), nums[k+1:]...)
			tolerance = false
			continue
		}
		i += 1
		k += 1
	}

	return true
}


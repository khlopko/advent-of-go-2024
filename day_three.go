package main

import (
	"bufio"
)

func dayThree(scanner *bufio.Scanner, part int) int {
	sum := 0
	stack1 := [3]byte{'m', 'u', 'l'}
	stack2 := [5]byte{'d', 'o', 'n', '\'', 't'}
	enabled := true

	for scanner.Scan() {
		text := scanner.Text()
		nums := [2]int{0, 0}

		p1 := 0
		p2 := 0
		i := 0
		for i < len(text) {
			if text[i] == stack1[p1] {
				p1 += 1
			}
			if text[i] == stack2[p2] {
				p2 += 1
			}
			i += 1
			if p2 == 2 && text[i] == '(' {
				p2 = 0
				enabled = true
			}
			if p2 == 5 && text[i] == '(' {
				p2 = 0
				enabled = false
			}
			if p1 == 3 {
				p1 = 0
				if enabled == false {
					continue
				}
				if text[i] != '(' {
					continue
				}
				i += 1

				nums[0] = 0
				for text[i] >= '0' && text[i] <= '9' {
					nums[0] *= 10
					nums[0] += int(text[i] - 48)
					i += 1
				}

				if text[i] != ',' {
					continue
				}
				i += 1

				nums[1] = 0
				for text[i] >= '0' && text[i] <= '9' {
					nums[1] *= 10
					nums[1] += int(text[i] - 48)
					i += 1
				}

				if text[i] != ')' {
					continue
				}
		
				sum += (nums[0] * nums[1])
			}
		}
	}
	return sum
}


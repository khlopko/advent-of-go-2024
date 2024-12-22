package main

import (
	"bufio"
	"slices"
)

func dayFive(scanner *bufio.Scanner, part int) int {
	scanningRules := true
	rules := make(map[int][]int)
	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanningRules = false
			continue
		}
		if scanningRules {
			pageAfter := parsePageNumber(line, [2]int{0, 1})
			pageBefore := parsePageNumber(line, [2]int{3, 4})
			rules[pageBefore] = append(rules[pageBefore], pageAfter)
		} else {
			i := 0
			update := []int{}
			for i < len(line) {
				update = append(update, parsePageNumber(line, [2]int{i, i + 1}))
				i += 3
			}
			updates = append(updates, update)
		}
	}

	midPages := []int{}
	for _, update := range updates {
		printedPages := make(map[int]bool)

		i := 0
		modified := false
		for i < len(update) {
			requiredPages := rules[update[i]]
			newPosition := -1
			for _, requiredPage := range requiredPages {
				if slices.Contains(update, requiredPage) == false {
					continue
				}

				if printedPages[requiredPage] {
					continue
				}

				k := i + 1
				for k < len(update) {
					if update[k] == requiredPage {
						newPosition = k
					}
					k += 1
				}
			}
			if newPosition != -1 {
				update = slices.Insert(update, newPosition + 1, update[i])
				update = slices.Delete(update, i, i+1)
				modified = true
			} else {
				printedPages[update[i]] = true
				i += 1
			}
		}
		if modified {
			midPages = append(midPages, update[len(update)/2])
		}
	}

	sum := 0
	for _, page := range midPages {
		sum += page
	}
	return sum
}

func parsePageNumber(line string, indexes [2]int) int {
	return int(line[indexes[0]]-48)*10 + int(line[indexes[1]]-48)
}

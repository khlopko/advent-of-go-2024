package main

import (
	"bufio"
)

func dayFour(scanner *bufio.Scanner, part int) int {
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	row := 0
	col := 0

	result := 0

	for row < len(lines) {
		line := lines[row]
		for col < len(line) {
			letter := line[col]
			if part == 1 {
				if letter == 'X' {
					result += lookupMas(lines, row, col)
				}
			}
			if part == 2 {
				if letter == 'A' {
					result += lookupX(lines, row, col)
				}
			}
			col += 1
		}
		col = 0
		row += 1
	}

	return result
}

var moves = [][]int{
	// up
	{-1, -1},
	{-1, 0},
	{-1, 1},

	// sides
	{0, -1},
	{0, 1},

	// down
	{1, -1},
	{1, 0},
	{1, 1},
}

var letters = [3]byte{'M', 'A', 'S'}

func lookupMas(lines []string, startRow int, startCol int) int {
	count := 0
	for _, move := range moves {
		row := startRow + move[0]
		col := startCol + move[1]
		idx := 0
		for idx >= 0 && idx < len(letters) && row >= 0 && row < len(lines) && col >= 0 && col < len(lines[row]) {
			if letters[idx] == lines[row][col] {
				idx += 1
			} else {
				idx = -1
			}
			row += move[0]
			col += move[1]
		}
		if idx == len(letters) {
			count += 1
		}
		idx = 0
	}
	return count
}

func lookupX(lines []string, startRow int, startCol int) int {
	row := startRow
	col := startCol
	if row > 0 && col > 0 && row < len(lines) - 1 && col < len(lines[row]) - 1 {
		topLeft := lines[row - 1][col - 1]
		topRight := lines[row - 1][col + 1]

		bottomLeft := lines[row + 1][col - 1]
		bottomRight := lines[row + 1][col + 1]

		rightDiag := (topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')
		leftDiag := (topRight == 'M' && bottomLeft == 'S') || (topRight == 'S' && bottomLeft == 'M')

		if rightDiag && leftDiag {
			return 1
		}
	}
	return 0
}

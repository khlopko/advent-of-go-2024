package main

import (
	"bufio"
)

func daySix(scanner *bufio.Scanner, part int, timing *Timing) int {
	timing.Start("day-6")

	labMap, mark, pos, direction := parseDaySixMap(scanner)

	visitedPos, _ := move(labMap, mark, pos, direction)

	if part == 1 {
		timing.Finish()
		return len(visitedPos)
	}

	count := 0
	for key := range visitedPos {
		labMap[key.y][key.x] = obstacleMark
		_, looped := move(labMap, mark, pos, direction)
		if looped {
			count += 1
		}
		labMap[key.y][key.x] = '.'
	}

	timing.Finish()
	return count
}

func move(labMap Map, mark GuardMark, startPos GuardPos, startDirection Direction) (map[GuardPos]Direction, bool) {
	visitedPos := make(map[GuardPos]Direction)
	looped := false
	pos := startPos
	direction := startDirection
	for {
		nextPos := GuardPos{pos.x+direction.h, pos.y+direction.v}
		if visitedPos[nextPos] == direction {
			looped = true
			break
		}
		if nextPos.y < 0 || nextPos.y >= len(labMap) || nextPos.x < 0 || nextPos.x >= len(labMap[nextPos.y]) {
			break
		}
		if labMap[nextPos.y][nextPos.x] == obstacleMark {
			mark = nextMark[mark]
			direction = guardMarkToDirection[mark]
		} else {
			pos = nextPos
			visitedPos[pos] = direction
		}
	}

	return visitedPos, looped
}

type Map = [][]byte

type GuardPos struct {
	x int
	y int
}

type Direction struct {
	v int
	h int
}

type GuardMark byte

var guardMarks = [4]GuardMark{'^', '>', 'v', '<'}
var obstacleMark byte = '#'
var guardMarkToDirection = map[GuardMark]Direction{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}
var nextMark = map[GuardMark]GuardMark {
	'^': '>',
	'>': 'v',
	'v': '<',
	'<': '^',
}

func parseDaySixMap(scanner *bufio.Scanner) (Map, GuardMark, GuardPos, Direction) {
	labMap := Map{}
	var initialMark GuardMark = 0
	startPos := GuardPos{-1,-1}
	direction := Direction{0, 0}

	row := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		labMap = append(labMap, []byte{})
		for col, b := range line {
			labMap[row] = append(labMap[row], b)
			for _, mark := range guardMarks {
				if mark == GuardMark(b) {
					initialMark = mark
					startPos = GuardPos{col, row}
					direction = guardMarkToDirection[GuardMark(b)]
				}
			}
		}
		row += 1
	}
	return labMap, initialMark, startPos, direction
}


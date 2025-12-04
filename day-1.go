package main

import (
	"fmt"
	"strconv"
)

type direction int

const (
	left direction = iota
	right
)

type dialer struct {
	direction direction
	steps     int
}

func decodeInstruction(c string) (dialer, error) {
	if len(c) < 2 {
		return dialer{}, fmt.Errorf("invalid command: too short")
	}

	var d dialer

	switch c[0] {
	case 'L':
		d.direction = left
	case 'R':
		d.direction = right
	default:
		return dialer{}, fmt.Errorf("invalid command: unknown direction '%c'", c[0])
	}

	rest := c[1:]
	for _, r := range rest {
		if r < '0' || r > '9' {
			return dialer{}, fmt.Errorf("invalid command: '%s' is not a number", rest)
		}
	}

	steps, err := strconv.Atoi(rest)
	check(err)
	d.steps = steps

	return d, nil
}

func day1Part1() {
	var start, count = 50, 0
	lineChan := readLineFromFile("./input/day-1.txt")

	for line := range lineChan {
		currDir, err := decodeInstruction(line)
		check(err)
		if currDir.direction == right {
			start = (start + currDir.steps) % 100
		} else {
			start = ((start-currDir.steps)%100 + 100) % 100
		}

		if start == 0 {
			count += 1
		}
	}
	fmt.Printf("part 1 count: %d\n", count)

}
func day1Part2() {
	start, count := 50, 0
	lineChan := readLineFromFile("./input/day-1.txt")
	for line := range lineChan {
		currDir, err := decodeInstruction(line)
		check(err)
		if currDir.direction == right {
			finish := (start + currDir.steps) % 100
			count += (start + currDir.steps) / 100
			start = finish
		} else {
			finish := ((start-currDir.steps)%100 + 100) % 100
			if start == 0 {
				count += currDir.steps / 100
			} else if currDir.steps >= start {
				count += 1 + (currDir.steps-start)/100
			}
			start = finish
		}
	}
	fmt.Printf("part 2 count: %d\n", count)
}

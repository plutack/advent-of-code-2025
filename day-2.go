package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isRepeatedPattern(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	firstHalf := s[:mid]
	secondHalf := s[mid:]
	return firstHalf == secondHalf
}

func day2Part1() {
	ranges, err := readSingleLineFile("./input/day-2.txt")
	check(err)
	total := 0
	for _, r := range ranges {
		items := strings.Split(r, "-")
		trimmedString := strings.TrimSpace(items[0])
		start, err := strconv.Atoi(trimmedString)
		check(err)
		trimmedString = strings.TrimSpace(items[1])
		end, err := strconv.Atoi(trimmedString)
		check(err)
		for n := start; n <= end; n++ {
			s := strconv.Itoa(n)
			if isRepeatedPattern(s) {
				total += n
			}
		}
	}
	fmt.Printf("total: %d\n", total)
}

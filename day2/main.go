package main

import (
	"strconv"
	"strings"
)

const RANGE_S = ','
const NUMBER_S = '-'

var text = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func part1() {
	var solution int = 0
	t := strings.Split(strings.Trim(text, "\n"), string(RANGE_S))
	for _, v := range t {
		parts := strings.Split(v, string(NUMBER_S))
		r1, _ := strconv.Atoi(parts[0])
		r2, _ := strconv.Atoi(parts[1])
		for i := r1; i <= r2; i++ {
			stringNumber := strconv.Itoa(i)
			windowSize := len(stringNumber) / 2
			if stringNumber[0:windowSize] == stringNumber[windowSize:] {
				solution = solution + i
			}
		}
	}
	println(solution)
}

func isValid(i int) bool {
	stringNumber := strconv.Itoa(i)
	nSize := len(stringNumber)
	for patternSize := 1; patternSize <= nSize/2; patternSize++ {
		if nSize%patternSize != 0 {
			continue
		}
		pattern := stringNumber[0:patternSize]
		p2 := strings.Repeat(pattern, nSize/patternSize)

		if p2 == stringNumber {
			return true
		}
	}
	return false
}

func part2() {
	var solution int = 0
	t := strings.Split(strings.Trim(text, "\n"), string(RANGE_S))
	for _, v := range t {
		parts := strings.Split(v, string(NUMBER_S))
		r1, _ := strconv.Atoi(parts[0])
		r2, _ := strconv.Atoi(parts[1])
		for i := r1; i <= r2; i++ {
			if isValid(i) {
				solution = solution + i
			}
		}
	}
	println(solution)
}

func main() {
	part1()
	part2()
}

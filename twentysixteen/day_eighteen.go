package twentysixteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func findSolutionDayEighteen(fp *bufio.Reader, rounds int) int {
	starting := utils.ReadSingleString(fp)

	safeCount := 0

	current := []byte(starting)
	next := make([]byte, len(current))

	for i := 0; i < rounds; i++ {
		for j := 0; j < len(next); j++ {
			var left, center, right bool
			if j-1 < 0 {
				left = false
			} else {
				left = current[j-1] == '^'
			}
			if j+1 == len(next) {
				right = false
			} else {
				right = current[j+1] == '^'
			}
			center = current[j] == '^'
			if current[j] == '^' {
				center = true
			} else {
				safeCount++
			}
			if (left && center && !right) || (!left && center && right) ||
				(left && !center && !right) || (!left && !center && right) {
				next[j] = '^'
			} else {
				next[j] = '.'
			}
		}
		tmp := current
		current = next
		next = tmp
	}

	return safeCount
}

func DayEighteenA(fp *bufio.Reader) string {
	solution := findSolutionDayEighteen(fp, 40)
	return strconv.Itoa(solution)
}

func DayEighteenB(fp *bufio.Reader) string {
	solution := findSolutionDayEighteen(fp, 400000)
	return strconv.Itoa(solution)
}

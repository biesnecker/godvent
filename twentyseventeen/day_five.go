package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func findSolutionDayFive(fp *bufio.Reader, partB bool) int {
	prog := utils.ReadOneIntegerPerLine(fp)

	idx := 0
	steps := 0
	for {
		if idx < 0 || idx >= len(prog) {
			return steps
		}
		steps++
		jump := prog[idx]
		if partB && prog[idx] > 2 {
			prog[idx]--
		} else {
			prog[idx]++
		}
		idx += jump
	}
}

func DayFiveA(fp *bufio.Reader) string {
	return strconv.Itoa(findSolutionDayFive(fp, false))
}

func DayFiveB(fp *bufio.Reader) string {
	return strconv.Itoa(findSolutionDayFive(fp, true))
}

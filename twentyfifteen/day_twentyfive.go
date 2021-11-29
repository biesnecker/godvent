package twentyfifteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayTwentyFive(fp *bufio.Reader) (int, int) {
	s := utils.ReadSingleString(fp)
	var tx, ty int
	fmt.Sscanf(
		s,
		"To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.",
		&tx, &ty)
	return tx, ty
}

func findSolutionDayTwentyFive(targetRow, targetColumn, seed int) int {
	row := 2
	col := 1
	for {
		crow := row
		ccol := col
		for crow > 0 {
			seed = (seed * 252533) % 33554393
			if crow == targetRow && ccol == targetColumn {
				return seed
			}
			crow -= 1
			ccol += 1
		}
		row += 1
	}
}

func DayTwentyFiveA(fp *bufio.Reader) string {
	targetRow, targetCol := readInputDayTwentyFive(fp)
	return strconv.Itoa(
		findSolutionDayTwentyFive(targetRow, targetCol, 20151125))
}

func DayTwentyFiveB(fp *bufio.Reader) string {
	return "There is no 2015 Day 25 Part B puzzle"
}

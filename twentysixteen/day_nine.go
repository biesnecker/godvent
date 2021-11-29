package twentysixteen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

func decompressDayNine(input string, recursive bool) int {
	total := 0
	for len(input) > 0 {
		nextParenOpen := strings.IndexByte(input, '(')
		if nextParenOpen == -1 {
			// No more groups to expand.
			total += len(input)
			break
		}
		total += nextParenOpen
		input = input[nextParenOpen:]

		var run, reps int
		fmt.Sscanf(input, "(%dx%d)", &run, &reps)

		nextParenClose := strings.IndexByte(input, ')')
		input = input[nextParenClose+1:]

		if recursive {
			total += reps * decompressDayNine(input[:run], recursive)
		} else {
			total += reps * run
		}
		input = input[run:]
	}
	return total
}

func DayNineA(fp *bufio.Reader) string {
	return strconv.Itoa(decompressDayNine(utils.ReadSingleString(fp), false))
}

func DayNineB(fp *bufio.Reader) string {
	return strconv.Itoa(decompressDayNine(utils.ReadSingleString(fp), true))

}

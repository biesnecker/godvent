package twentyseventeen

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

func findSolutionDayFour(fp *bufio.Reader, shouldSort bool) int {
	total := 0
	utils.ReadStrings(fp, func(s string) {
		m := make(map[string]struct{})
		parts := strings.Fields(s)
		for i := range parts {
			if shouldSort {
				m[utils.SortString(parts[i])] = struct{}{}
			} else {
				m[parts[i]] = struct{}{}
			}
		}
		if len(m) == len(parts) {
			total++
		}
	})
	return total
}

func DayFourA(fp *bufio.Reader) string {
	return strconv.Itoa(findSolutionDayFour(fp, false))
}

func DayFourB(fp *bufio.Reader) string {
	return strconv.Itoa(findSolutionDayFour(fp, true))
}

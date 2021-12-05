package twentynineteen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

func getRangeFromInput(fp *bufio.Reader) (int, int) {
	s := utils.ReadSingleString(fp)
	s = strings.ReplaceAll(s, "-", " ")
	var a, b int
	fmt.Sscanf(s, "%d %d", &a, &b)
	return a, b
}

func isValidDay4a(n int) bool {
	hasPair := false
	s := strconv.Itoa(n)
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
		if s[i] == s[i-1] {
			hasPair = true
		}
	}
	return hasPair
}

func isValidDay4b(n int) bool {
	consecutiveCount := 0
	hasPair := false
	s := strconv.Itoa(n)
	for i := 1; i < len(s); i++ {
		if s[i] < s[i-1] {
			return false
		}
		if s[i] == s[i-1] {
			consecutiveCount++
		} else {
			if consecutiveCount == 1 {
				hasPair = true
			}
			consecutiveCount = 0
		}
	}
	return consecutiveCount == 1 || hasPair
}

func DayFourA(fp *bufio.Reader) string {
	start, end := getRangeFromInput(fp)
	count := 0
	for i := start; i <= end; i++ {
		if isValidDay4a(i) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func DayFourB(fp *bufio.Reader) string {
	start, end := getRangeFromInput(fp)
	count := 0
	for i := start; i <= end; i++ {
		if isValidDay4b(i) {
			count++
		}
	}
	return strconv.Itoa(count)
}

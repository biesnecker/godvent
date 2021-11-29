package twentyfifteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func isGoodStringFiveA(s *string) bool {
	slen := len(*s)
	vowels := 0
	hasDouble := false
	for i := 0; i < slen; i++ {
		c := (*s)[i]
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			vowels++
		}
		if i < slen-1 {
			d := (*s)[i+1]
			if !hasDouble {
				hasDouble = c == d
			}
			if (c == 'a' && d == 'b') ||
				(c == 'c' && d == 'd') ||
				(c == 'p' && d == 'q') ||
				(c == 'x' && d == 'y') {
				return false
			}
		}
	}
	return vowels >= 3 && hasDouble
}

func isGoodStringFiveB(s *string) bool {
	slen := len(*s)
	hasPair := false
	hasAxA := false
	for i := 0; i < slen-2; i++ {
		if (*s)[i] == (*s)[i+2] {
			hasAxA = true
			break
		}
	}

	for i := 0; i < slen-3; i++ {
		fst := (*s)[i]
		snd := (*s)[i+1]
		for j := i + 2; j < slen-1; j++ {
			if fst == (*s)[j] && snd == (*s)[j+1] {
				hasPair = true
				break
			}
		}
		if hasPair {
			break
		}
	}

	return hasPair && hasAxA
}

func DayFiveA(fp *bufio.Reader) string {
	total := 0
	utils.ReadStrings(fp, func(s string) {
		if isGoodStringFiveA(&s) {
			total++
		}
	})
	return strconv.Itoa(total)
}

func DayFiveB(fp *bufio.Reader) string {
	total := 0
	utils.ReadStrings(fp, func(s string) {
		if isGoodStringFiveB(&s) {
			total++
		}
	})
	return strconv.Itoa(total)
}

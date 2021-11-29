package twentyfifteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DayEightA(fp *bufio.Reader) string {
	totalChars := 0
	stringChars := 0
	inEscape := false
	utils.ReadStrings(fp, func(s string) {
		s = s[1 : len(s)-1]
		slen := len(s)
		totalChars += 2 // for the beginning and ending quotes
		idx := 0
		for {
			if idx >= slen {
				break
			}
			c := s[idx]
			totalChars++
			if inEscape {
				if c == 'x' {
					idx += 3
					totalChars += 2
					stringChars++
				} else {
					idx++
					stringChars++
				}
				inEscape = false
			} else {
				if c == '\\' {
					inEscape = true
				} else {
					stringChars++
				}
				idx++
			}
		}
	})
	return strconv.Itoa(totalChars - stringChars)
}

func DayEightB(fp *bufio.Reader) string {
	literalChars := 0
	encodedChars := 0
	utils.ReadStrings(fp, func(s string) {
		literalChars += len(s)
		encodedChars += 2 // for the quotes
		for _, c := range s {
			if c == '\\' || c == '"' {
				encodedChars++
			}
			encodedChars++
		}
	})
	return strconv.Itoa(encodedChars - literalChars)
}

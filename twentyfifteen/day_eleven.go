package twentyfifteen

import (
	"bufio"

	"github.com/biesnecker/godvent/utils"
)

func isValidPasswordDayEleven(pw []byte) bool {
	pwlen := len(pw)
	idx := 0
	hasSeq := false
	firstPair := byte(0)
	hasPair := false

	for {
		if idx >= pwlen {
			break
		}
		c := pw[idx]
		if c == 'i' || c == 'o' || c == 'l' {
			return false
		}
		if !hasSeq &&
			(idx < pwlen-2) &&
			(c == pw[idx+1]-1) &&
			(c == pw[idx+2]-2) {
			hasSeq = true
		}
		if !hasPair && (idx < pwlen-1) && (c == pw[idx+1]) && (c != firstPair) {
			if firstPair == 0 {
				firstPair = c
			} else {
				hasPair = true
			}
		}
		idx++
	}
	return hasSeq && hasPair
}

func iterateUntilValidPassword(starting string) string {
	password := []byte(starting)
	rightIdx := len(password) - 1
	currentIdx := rightIdx
	for {
		nextByte := password[currentIdx] + 1
		if nextByte == 'i' || nextByte == 'l' || nextByte == 'o' {
			nextByte++
		}
		if nextByte > 'z' {
			for i := currentIdx; i <= rightIdx; i++ {
				password[i] = 'a'
			}
			currentIdx -= 1
		} else {
			password[currentIdx] = nextByte
			currentIdx = rightIdx
		}
		if isValidPasswordDayEleven(password) {
			return string(password)
		}
	}
}

func DayElevenA(fp *bufio.Reader) string {
	return iterateUntilValidPassword(utils.ReadSingleString(fp))
}

func DayElevenB(fp *bufio.Reader) string {
	// Start from the previous found password, hxbxxyzz
	return iterateUntilValidPassword("hxbxxyzz")

}

package twentysixteen

import (
	"bufio"
	"container/list"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type entryD14 struct {
	idx         int
	three, five byte
}

func hasThree(s string) (bool, byte) {
	for i := range s[:len(s)-2] {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return true, s[i]
		}
	}
	return false, 0
}

func hasFive(s string) (bool, byte) {
	for i := range s[:len(s)-4] {
		if s[i] == s[i+1] && s[i] == s[i+2] && s[i] == s[i+3] && s[i] == s[i+4] {
			return true, s[i]
		}
	}
	return false, 0
}

func findSolutionDayFourteen(salt string, stretch bool) int {

	idx := 0
	found := 0
	foundIdx := 0

	entries := list.New()

	for found < 64 {
		hash := utils.GetMD5StringOfString(fmt.Sprintf("%s%d", salt, idx))
		if stretch {
			for i := 0; i < 2016; i++ {
				hash = utils.GetMD5StringOfString(hash)
			}
		}

		hThree, three := hasThree(hash)
		hFive, five := hasFive(hash)
		if hThree || hFive {
			entry := entryD14{idx: idx, three: three, five: five}
			entries.PushBack(&entry)

			for entries.Len() > 1 {
				headEntry := entries.Front()
				head := headEntry.Value.(*entryD14)
				if head.idx+1000 > idx {
					// Don't bother processing anything until we're sure that
					// there's enough to remove nodes.
					break
				}

				nextEntry := headEntry.Next()
				for nextEntry != nil {
					next := nextEntry.Value.(*entryD14)
					if head.three == next.five && head.idx+1000 >= next.idx && found < 64 {
						found++
						foundIdx = head.idx
					}
					nextEntry = nextEntry.Next()
				}
				entries.Remove(headEntry)
			}
		}
		idx++
	}
	return foundIdx
}

func DayFourteenA(fp *bufio.Reader) string {
	salt := utils.ReadSingleString(fp)
	solution := findSolutionDayFourteen(salt, false)
	return strconv.Itoa(solution)
}

func DayFourteenB(fp *bufio.Reader) string {
	salt := utils.ReadSingleString(fp)
	solution := findSolutionDayFourteen(salt, true)
	return strconv.Itoa(solution)
}

package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type streamGroup struct {
	value, depth int
	prev         *streamGroup
}

func findSolutionDayNine(fp *bufio.Reader) (int, int) {
	h := streamGroup{value: 0, depth: 0, prev: nil}
	head := &h
	skipNext := false
	inGarbage := false
	garbageCount := 0
	utils.ReadChars(fp, func(c rune, _ int) bool {
		if skipNext {
			skipNext = false
			return true
		}
		switch c {
		case '<':
			if !inGarbage {
				inGarbage = true
			} else {
				garbageCount++
			}
		case '>':
			if inGarbage {
				inGarbage = false
			}
		case '{':
			if !inGarbage {
				nsg := &streamGroup{
					value: head.depth + 1,
					depth: head.depth + 1,
					prev:  head}
				head = nsg
			} else {
				garbageCount++
			}
		case '}':
			if !inGarbage {
				headValue := head.value
				head = head.prev
				head.value += headValue
			} else {
				garbageCount++
			}
		case '!':
			skipNext = inGarbage
		default:
			skipNext = false
			if inGarbage {
				garbageCount++
			}
		}

		return true
	})
	return h.value, garbageCount
}

func DayNineA(fp *bufio.Reader) string {
	total, _ := findSolutionDayNine(fp)
	return strconv.Itoa(total)
}

func DayNineB(fp *bufio.Reader) string {
	_, garbageCount := findSolutionDayNine(fp)
	return strconv.Itoa(garbageCount)
}

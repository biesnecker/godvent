package twentyfifteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DayOneA(fp *bufio.Reader) string {
	s := 0
	utils.ReadChars(fp, func(c rune, _ int) bool {
		switch c {
		case '(':
			s++
		case ')':
			s--
		default:
			return false
		}
		return true
	})
	return strconv.Itoa(s)
}

func DayOneB(fp *bufio.Reader) string {
	var s struct{ floor, idx int }
	utils.ReadChars(fp, func(c rune, _ int) bool {
		switch c {
		case '(':
			s.floor++
		case ')':
			s.floor--
		default:
			return false
		}
		s.idx++
		return s.floor != -1
	})
	return strconv.Itoa(s.idx)
}

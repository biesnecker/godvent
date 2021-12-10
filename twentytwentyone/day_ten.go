package twentytwentyone

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayTen(fp *bufio.Reader) {

}

func DayTenA(fp *bufio.Reader) string {
	var stack [10000]byte
	var tos int

	counts := make(map[byte]int)

	utils.ReadStrings(fp, func(s string) {
		tos = 0
	lineiter:
		for i := range s {
			switch s[i] {
			case '(', '[', '<', '{':
				stack[tos] = s[i]
				tos++
			default:
				tos--
				switch stack[tos] {
				case '(':
					if s[i] != ')' {
						counts[s[i]]++
						break lineiter
					}
				case '[':
					if s[i] != ']' {
						counts[s[i]]++
						break lineiter
					}
				case '<':
					if s[i] != '>' {

						counts[s[i]]++
						break lineiter
					}
				case '{':
					if s[i] != '}' {

						counts[s[i]]++
						break lineiter
					}
				}

			}
		}
	})
	total := 0
	for k, v := range counts {
		switch k {
		case ')':
			total += (3 * v)
		case ']':
			total += (57 * v)
		case '}':
			total += (1197 * v)
		case '>':
			total += (25137 * v)
		}
	}
	return strconv.Itoa(total)
}

func DayTenB(fp *bufio.Reader) string {
	var stack [10000]byte
	var tos int

	scores := make([]int, 0, 100)

	utils.ReadStrings(fp, func(s string) {
		tos = 0
		isValid := true
	lineiter:
		for i := range s {
			switch s[i] {
			case '(', '[', '<', '{':
				stack[tos] = s[i]
				tos++
			default:
				tos--
				switch stack[tos] {
				case '(':
					if s[i] != ')' {
						isValid = false
						break lineiter
					}
				case '[':
					if s[i] != ']' {
						isValid = false
						break lineiter
					}
				case '<':
					if s[i] != '>' {
						isValid = false
						break lineiter
					}
				case '{':
					if s[i] != '}' {
						isValid = false
						break lineiter
					}
				}

			}
		}
		if isValid && tos > 0 {
			score := 0
			for i := tos - 1; i >= 0; i-- {
				score *= 5
				switch stack[i] {
				case '(':
					score += 1
				case '[':
					score += 2
				case '{':
					score += 3
				case '<':
					score += 4
				}
			}
			scores = append(scores, score)
		}
	})

	sort.Sort(sort.IntSlice(scores))

	return strconv.Itoa(scores[len(scores)/2])
}

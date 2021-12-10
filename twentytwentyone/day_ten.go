package twentytwentyone

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func bracketMatch(left, right byte) bool {
	return (left == '(' && right == ')') ||
		(left == '[' && right == ']') ||
		(left == '{' && right == '}') ||
		(left == '<' && right == '>')
}

var bracketScores map[byte]int = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
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
				if !bracketMatch(stack[tos], s[i]) {
					counts[s[i]]++
					break lineiter
				}
			}
		}
	})
	total := 0
	for k, v := range counts {
		total += bracketScores[k] * v
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
				if !bracketMatch(stack[tos], s[i]) {
					isValid = false
					break lineiter
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

	sort.Ints(scores)

	return strconv.Itoa(scores[len(scores)/2])
}

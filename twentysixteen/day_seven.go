package twentysixteen

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

func parseInputStringDaySeven(s string) ([]string, []string) {
	var outside, inside []string

	for len(s) > 0 {
		nextOpen := strings.IndexByte(s, '[')
		if nextOpen == -1 {
			// There are no more insides, what's left is outside.
			outside = append(outside, s)
			break
		} else {
			// Push everything before that to outside.
			if nextOpen > 0 {
				outside = append(outside, s[:nextOpen])
			}
			s = s[nextOpen+1:]
			nextClose := strings.IndexByte(s, ']')
			if nextClose == -1 {
				log.Fatal("Missing close bracket")
			}
			inside = append(inside, s[:nextClose])
			s = s[nextClose+1:]
		}
	}
	return outside, inside
}

func hasABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
			return true
		}
	}
	return false
}

type abamatch struct {
	a, b byte
}

func hasABA(s string) []abamatch {
	var matches []abamatch
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] && s[i] != s[i+1] {
			matches = append(matches, abamatch{a: s[i], b: s[i+1]})
		}
	}
	return matches
}

func DaySevenA(fp *bufio.Reader) string {
	total := 0
	utils.ReadStrings(fp, func(s string) {
		outside, inside := parseInputStringDaySeven(s)
		hasOutside := false
		hasInside := false
		for _, out := range outside {
			if hasABBA(out) {
				hasOutside = true
				break
			}
		}
		for _, in := range inside {
			if hasABBA(in) {
				hasInside = true
				break
			}
		}
		if hasOutside && !hasInside {
			total += 1
		}
	})
	return strconv.Itoa(total)
}

func DaySevenB(fp *bufio.Reader) string {
	total := 0
	utils.ReadStrings(fp, func(s string) {
		outside, inside := parseInputStringDaySeven(s)
		var outsideMatches []abamatch
		var insideMatches []abamatch

		for _, out := range outside {
			m := hasABA(out)
			outsideMatches = append(outsideMatches, m...)
		}
		for _, in := range inside {
			m := hasABA(in)
			insideMatches = append(insideMatches, m...)
		}

		for _, i := range outsideMatches {
			finished := false
			for _, j := range insideMatches {
				if i.a == j.b && j.a == i.b {
					total++
					finished = true
					break
				}
			}
			if finished {
				break
			}
		}
	})
	return strconv.Itoa(total)
}

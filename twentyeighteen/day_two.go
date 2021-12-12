package twentyeighteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func DayTwoA(fp *bufio.Reader) string {
	var twos, threes int
	utils.ReadStrings(fp, func(s string) {
		counts := make(map[byte]int)
		for i := range s {
			counts[s[i]]++
		}
		var hasTwos, hasThrees bool
		for _, c := range counts {
			if c == 2 {
				hasTwos = true
			} else if c == 3 {
				hasThrees = true
			}
		}
		if hasTwos {
			twos++
		}
		if hasThrees {
			threes++
		}
	})
	return strconv.Itoa(twos * threes)
}

func DayTwoB(fp *bufio.Reader) string {
	lines := utils.ReadStringsAsSlice(fp)
	for i := range lines {
	nextline:
		for j := range lines {
			if i == j {
				continue
			}

			diffpos := -1
			for k := 0; k < len(lines[i]); k++ {
				if lines[i][k] != lines[j][k] {
					if diffpos >= 0 {
						continue nextline
					}
					diffpos = k
				}
			}
			res := make([]byte, 0, len(lines[i]))
			for k := 0; k < len(lines[i]); k++ {
				if k == diffpos {
					continue
				}
				res = append(res, lines[i][k])
			}
			return string(res)
		}
	}
	return ""
}

package twentyseventeen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type inputDay24 struct {
	a, b int
}

type stepResult24 struct {
	absoluteMax int
	longestLen  int
	longestMax  int
}

func readInputDayTwentyFour(fp *bufio.Reader) []inputDay24 {
	res := make([]inputDay24, 0, 60)
	utils.ReadStrings(fp, func(s string) {
		i := inputDay24{}
		fmt.Sscanf(s, "%d/%d", &i.a, &i.b)
		res = append(res, i)
	})
	return res
}

func makeBridgeD24(elems []inputDay24, used uint64, needed int, current stepResult24) stepResult24 {
	best := current
elemLoop:
	for i := range elems {
		mask := uint64(1) << i
		if used&mask > 0 {
			continue elemLoop
		}
		var otherside int
		if elems[i].a == needed {
			otherside = elems[i].b
		} else if elems[i].b == needed {
			otherside = elems[i].a
		} else {
			continue elemLoop
		}

		elemScore := elems[i].a + elems[i].b

		next := current
		next.absoluteMax += elemScore
		next.longestLen++
		next.longestMax += elemScore

		res := makeBridgeD24(elems, used|mask, otherside, next)
		if res.absoluteMax > best.absoluteMax {
			best.absoluteMax = res.absoluteMax
		}
		if res.longestLen > best.longestLen {
			best.longestLen = res.longestLen
			best.longestMax = res.longestMax
		} else if res.longestLen == best.longestLen && best.longestMax < res.longestMax {
			best.longestMax = res.longestMax
		}
	}
	return best
}

func DayTwentyFourA(fp *bufio.Reader) string {
	input := readInputDayTwentyFour(fp)
	solution := makeBridgeD24(input, 0, 0, stepResult24{})
	return strconv.Itoa(solution.absoluteMax)
}

func DayTwentyFourB(fp *bufio.Reader) string {
	input := readInputDayTwentyFour(fp)
	solution := makeBridgeD24(input, 0, 0, stepResult24{})
	return strconv.Itoa(solution.longestMax)
}

package twentysixteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func isTriangleDayThree(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

func DayThreeA(fp *bufio.Reader) string {
	count := 0
	utils.ReadStrings(fp, func(s string) {
		var a, b, c int
		fmt.Sscanf(s, "%d %d %d", &a, &b, &c)
		if isTriangleDayThree(a, b, c) {
			count++
		}
	})
	return strconv.Itoa(count)
}

func DayThreeB(fp *bufio.Reader) string {
	count := 0

	colA := make([]int, 0, 3)
	colB := make([]int, 0, 3)
	colC := make([]int, 0, 3)
	utils.ReadStrings(fp, func(s string) {
		var a, b, c int
		fmt.Sscanf(s, "%d %d %d", &a, &b, &c)

		colA = append(colA, a)
		colB = append(colB, b)
		colC = append(colC, c)

		if len(colA) == 3 {
			if isTriangleDayThree(colA[0], colA[1], colA[2]) {
				count++
			}
			if isTriangleDayThree(colB[0], colB[1], colB[2]) {
				count++
			}
			if isTriangleDayThree(colC[0], colC[1], colC[2]) {
				count++
			}
			colA = make([]int, 0, 3)
			colB = make([]int, 0, 3)
			colC = make([]int, 0, 3)
		}
	})
	return strconv.Itoa(count)
}

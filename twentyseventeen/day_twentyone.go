package twentyseventeen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

func getFormatString(s string) string {
	return strings.ReplaceAll(s, "/", "")
}

func getSideLengthFromBytes(m []byte) int {
	byteLength := len(m)
	if byteLength == 4 {
		return 2
	} else if byteLength == 9 {
		return 3
	} else if byteLength == 16 {
		return 4
	} else {
		panic("Unknown byte size")
	}
}

func flipBytes(orig string) string {
	m := make([]byte, len(orig))
	copy(m, orig)
	n := getSideLengthFromBytes(m)
	for y := 0; y < n/2; y++ {
		for x := 0; x < n; x++ {
			m[x*n+y], m[x*n+(n-1-y)] = m[x*n+(n-1-y)], m[x*n+y]
		}
	}
	return string(m)
}

func rotateBytes(orig string) string {
	m := make([]byte, len(orig))
	copy(m, orig)
	n := getSideLengthFromBytes(m)

	var a, b, c, d byte

	for i := 0; i <= n/2-1; i++ {
		for j := 0; j <= n-(2*i)-2; j++ {
			a = m[n*(i+j)+i]
			b = m[n*(n-1-i)+i+j]
			c = m[n*(n-1-i-j)+(n-1-i)]
			d = m[n*i+(n-1-i-j)]

			m[n*(i+j)+i] = b
			m[n*(n-1-i)+i+j] = c
			m[n*(n-1-i-j)+(n-1-i)] = d
			m[n*i+(n-1-i-j)] = a
		}
	}
	return string(m)
}

func getSubgrids(m string) []string {
	res := [][]byte{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}

	if len(m) != 16 {
		panic("Must be 4x4")
	}
	res[0][0] = m[0]
	res[0][1] = m[1]
	res[0][2] = m[4]
	res[0][3] = m[5]
	res[1][0] = m[2]
	res[1][1] = m[3]
	res[1][2] = m[6]
	res[1][3] = m[7]
	res[2][0] = m[8]
	res[2][1] = m[9]
	res[2][2] = m[0xC]
	res[2][3] = m[0xD]
	res[3][0] = m[0xA]
	res[3][1] = m[0xB]
	res[3][2] = m[0xE]
	res[3][3] = m[0xF]

	return []string{
		string(res[0]), string(res[1]),
		string(res[2]), string(res[3])}
}

func readInputDayTwentyOne(fp *bufio.Reader) map[string]string {
	replacements := make(map[string]string)
	utils.ReadStrings(fp, func(s string) {
		parts := strings.Split(s, " => ")
		left := getFormatString(parts[0])
		right := getFormatString(parts[1])

		for i := 0; i < 4; i++ {
			replacements[left] = right
			replacements[flipBytes(left)] = right
			left = rotateBytes(left)
		}
	})
	return replacements
}

func stepD21(replacements map[string]string, pattern string, step, maxStep int) int {
	fmt.Println(strings.Repeat("\t", step), len(pattern))
	if step == maxStep {
		count := 0
		for _, c := range pattern {
			if c == '#' {
				count++
			}
		}
		fmt.Println(strings.Repeat("\t", step), "total:", count)
		return count
	}

	if replacement, found := replacements[pattern]; !found {
		panic(fmt.Sprintf("Could not find replacement for %s\n", pattern))
	} else if len(replacement) == 16 {
		fmt.Println(strings.Repeat("\t", step), len(replacement))
		total := 0
		for _, subgrid := range getSubgrids(replacement) {
			total += stepD21(replacements, subgrid, step+1, maxStep)
		}
		fmt.Println(strings.Repeat("\t", step), "total:", total)
		return total
	} else {
		fmt.Println(strings.Repeat("\t", step), len(replacement))
		total := stepD21(replacements, replacement, step+1, maxStep)
		fmt.Println(strings.Repeat("\t", step), "total:", total)
		return total
	}
}

func DayTwentyOneA(fp *bufio.Reader) string {
	replacements := readInputDayTwentyOne(fp)

	initial := ".#...####"

	solution := stepD21(replacements, initial, 0, 3)
	return strconv.Itoa(solution)
}

func DayTwentyOneB(fp *bufio.Reader) string {
	return ""
}

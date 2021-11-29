package twentysixteen

import (
	"bufio"

	"github.com/biesnecker/godvent/utils"
)

func stepDay16(code []bool, goal int) []bool {
	res := make([]bool, goal)
	finished := false

	copy(res, code)

	currentLen := len(code)
finished:
	for !finished {
		for i := 0; i < currentLen; i++ {
			idx := currentLen + i + 1
			if idx >= goal {
				break finished
			}
			res[idx] = !res[currentLen-1-i]
		}
		currentLen = currentLen*2 + 1
	}
	return res
}

func calculateChecksum(code []bool, goal int) string {
	checksum := code[:goal]
	for len(checksum)%2 == 0 {
		for i := 0; i < len(checksum); i += 2 {
			checksum[i/2] = checksum[i] == checksum[i+1]
		}
		checksum = checksum[:len(checksum)/2]
	}
	result := make([]byte, len(checksum))
	for i, val := range checksum {
		if val {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}

func readInputDaySixteen(fp *bufio.Reader) []bool {
	var res []bool
	for _, b := range utils.ReadSingleString(fp) {
		if b == '1' {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}

func findSolutionDaySixteen(fp *bufio.Reader, goal int) string {
	code := readInputDaySixteen(fp)
	code = stepDay16(code, goal)
	return calculateChecksum(code, goal)
}

func DaySixteenA(fp *bufio.Reader) string {
	return findSolutionDaySixteen(fp, 272)
}

func DaySixteenB(fp *bufio.Reader) string {
	return findSolutionDaySixteen(fp, 35651584)
}

package twentytwentyone

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayThree(fp *bufio.Reader) {
	utils.ReadStrings(fp, func(s string) {

	})
}

func DayThreeA(fp *bufio.Reader) string {
	input := utils.ReadStringsAsSlice(fp)

	ones := make([]int, 12)
	zeros := make([]int, 12)

	for _, s := range input {
		for i, c := range s {
			if c == '1' {
				ones[i]++
			} else {
				zeros[i]++
			}
		}
	}
	var gammaB []byte
	var epsilonB []byte
	for i := 0; i < len(ones); i++ {
		if ones[i] > zeros[i] {
			gammaB = append(gammaB, '1')
			epsilonB = append(epsilonB, '0')
		} else {
			gammaB = append(gammaB, '0')
			epsilonB = append(epsilonB, '1')
		}
	}
	gamma, _ := strconv.ParseInt(string(gammaB), 2, 64)
	epsilon, _ := strconv.ParseInt(string(epsilonB), 2, 64)
	return strconv.Itoa(int(gamma) * int(epsilon))
}

func findCandidate(cs []string, most bool) string {
	i := 0
	for len(cs) > 1 {
		var zeros, ones int
		for _, c := range cs {
			if c[i] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		var matchbit byte
		if most {
			if zeros > ones {
				matchbit = '0'
			} else {
				matchbit = '1'
			}
		} else {
			if zeros > ones {
				matchbit = '1'
			} else {
				matchbit = '0'
			}
		}
		var matches []string
		for _, c := range cs {
			if c[i] == matchbit {
				matches = append(matches, c)
			}
		}
		cs = matches
		i++
	}
	return cs[0]
}

func DayThreeB(fp *bufio.Reader) string {
	oxygenCandidates := utils.ReadStringsAsSlice(fp)

	co2Candidates := make([]string, len(oxygenCandidates))
	copy(co2Candidates, oxygenCandidates)

	oxygen := findCandidate(oxygenCandidates, true)
	co2 := findCandidate(co2Candidates, false)

	oi, _ := strconv.ParseInt(oxygen, 2, 64)
	ci, _ := strconv.ParseInt(co2, 2, 64)

	return strconv.Itoa(int(oi) * int(ci))
}

package twentytwentyone

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func readInputD14(fp *bufio.Reader) (string, map[string]byte) {
	m := make(map[string]byte)
	var res string
	first := true
	utils.ReadStrings(fp, func(s string) {
		if first {
			res = s
			first = false
		} else if len(s) == 0 {
			return
		} else {
			var left, right string
			fmt.Sscanf(s, "%s -> %s", &left, &right)
			m[left] = right[0]
		}
	})
	return res, m
}

func getAnswerDay14(original string, m map[string]byte, rounds int) int {
	counts := make(map[string]int)

	for i := 1; i < len(original); i++ {
		key := string([]byte{original[i-1], original[i]})
		counts[key]++
	}

	for i := 0; i < rounds; i++ {
		newCounts := make(map[string]int)
		for k, v := range counts {
			mid := m[k]
			keyA := string([]byte{k[0], mid})
			keyB := string([]byte{mid, k[1]})
			newCounts[keyA] += v
			newCounts[keyB] += v
		}
		counts = newCounts
	}

	byteCounts := make(map[byte]int)
	for k, v := range counts {
		byteCounts[k[0]] += v
	}
	// Because we only count the first byte of each pair we need to add the
	// last byte of the original, which will always be at the end of the last
	// pair no matter how large the string becomes.
	byteCounts[original[len(original)-1]]++

	c := make([]int, 0, len(counts))

	for _, v := range byteCounts {
		c = append(c, v)
	}

	sort.Ints(c)

	return c[len(c)-1] - c[0]
}

func DayFourteenA(fp *bufio.Reader) string {
	res, m := readInputD14(fp)
	return strconv.Itoa(getAnswerDay14(res, m, 10))
}

func DayFourteenB(fp *bufio.Reader) string {
	res, m := readInputD14(fp)
	return strconv.Itoa(getAnswerDay14(res, m, 40))
}

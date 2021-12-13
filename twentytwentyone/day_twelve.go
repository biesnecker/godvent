package twentytwentyone

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

type nodeD12 struct {
	label    string
	next     []*nodeD12
	smallIdx int
}

func readInputDayTwelve(fp *bufio.Reader) map[string]*nodeD12 {
	res := make(map[string]*nodeD12)
	smallIdx := 1

	getNode := func(label string) *nodeD12 {
		if n, ok := res[label]; ok {
			return n
		} else {
			sidx := -1
			if unicode.IsLower(rune(label[0])) {
				if label == "start" {
					sidx = 0
				} else {
					sidx = smallIdx
					smallIdx++
				}
			}
			n := nodeD12{label: label, smallIdx: sidx}
			res[label] = &n
			return &n
		}
	}

	utils.ReadStrings(fp, func(s string) {
		s = strings.ReplaceAll(s, "-", " ")
		var left, right string
		fmt.Sscanf(s, "%s %s", &left, &right)
		leftN := getNode(left)
		rightN := getNode(right)

		leftN.next = append(leftN.next, rightN)
		rightN.next = append(rightN.next, leftN)
	})

	return res
}

func getPathD12(n *nodeD12, seen uint16, partTwo, doneTwice bool, count *int) {
	for _, nn := range n.next {
		if nn.label == "end" {
			*count++
			continue
		}

		newDoneTwice := doneTwice
		if nn.smallIdx >= 0 && seen&(uint16(1)<<uint16(nn.smallIdx)) > 0 {
			if !partTwo || nn.smallIdx == 0 {
				continue
			} else if partTwo && doneTwice {
				continue
			} else {
				newDoneTwice = true
			}
		}

		newseen := seen
		if nn.smallIdx >= 0 {
			newseen |= uint16(1) << uint16(nn.smallIdx)
		}
		getPathD12(nn, newseen, partTwo, newDoneTwice, count)
	}

}

func DayTwelveA(fp *bufio.Reader) string {
	input := readInputDayTwelve(fp)
	count := 0
	getPathD12(input["start"], 1, false, false, &count)
	return strconv.Itoa(count)
}

func DayTwelveB(fp *bufio.Reader) string {
	input := readInputDayTwelve(fp)
	count := 0
	getPathD12(input["start"], 1, true, false, &count)
	return strconv.Itoa(count)
}

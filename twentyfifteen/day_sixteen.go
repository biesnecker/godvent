package twentyfifteen

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

func getMFCSAMOutput() map[string]int {
	return map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
}

func handleInputDaySixteen(fp *bufio.Reader, isMatch func(map[string]int, string, int) bool) int {
	sueId := 0
	output := getMFCSAMOutput()
	utils.ReadStringsUntilBreak(fp, true, func(s string) bool {
		var id, attrOne, attrTwo, attrThree int
		var attrOneName, attrTwoName, attrThreeName string

		if _, err := fmt.Sscanf(s, "Sue %d: %s %d, %s %d, %s %d",
			&id, &attrOneName, &attrOne, &attrTwoName, &attrTwo,
			&attrThreeName, &attrThree,
		); err != nil {
			log.Fatal(err)
		}

		type attr struct {
			name  string
			value int
		}

		attrs := []attr{
			{strings.TrimSuffix(attrOneName, ":"), attrOne},
			{strings.TrimSuffix(attrTwoName, ":"), attrTwo},
			{strings.TrimSuffix(attrThreeName, ":"), attrThree},
		}

		for _, a := range attrs {
			if !isMatch(output, a.name, a.value) {
				return true
			}
		}

		sueId = id
		return false
	})
	return sueId
}

func DaySixteenA(fp *bufio.Reader) string {
	sueId := handleInputDaySixteen(fp, func(attrs map[string]int, name string, value int) bool {
		return attrs[name] == value
	})
	return strconv.Itoa(sueId)
}

func DaySixteenB(fp *bufio.Reader) string {
	sueId := handleInputDaySixteen(fp, func(attrs map[string]int, name string, value int) bool {
		attrValue := attrs[name]
		if name == "cats" || name == "trees" {
			return value > attrValue
		} else if name == "pomeranians" || name == "goldfish" {
			return value < attrValue
		} else {
			return value == attrValue
		}
	})
	return strconv.Itoa(sueId)
}

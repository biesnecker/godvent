package twentytwentyone

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type inputDayTwo struct {
	command string
	arg     int
}

func readInputDayTwo(fp *bufio.Reader) (res []inputDayTwo) {
	utils.ReadStrings(fp, func(s string) {
		var c string
		var a int
		fmt.Sscanf(s, "%s %d", &c, &a)
		res = append(res, inputDayTwo{c, a})
	})
	return
}

func DayTwoA(fp *bufio.Reader) string {
	input := readInputDayTwo(fp)
	depth := 0
	pos := 0
	for _, c := range input {
		switch c.command {
		case "forward":
			pos += c.arg
		case "down":
			depth += c.arg
		case "up":
			depth -= c.arg
		}
	}
	return strconv.Itoa(depth * pos)
}

func DayTwoB(fp *bufio.Reader) string {
	input := readInputDayTwo(fp)
	depth := 0
	pos := 0
	aim := 0
	for _, c := range input {
		switch c.command {
		case "forward":
			pos += c.arg
			depth += aim * c.arg
		case "down":
			aim += c.arg
		case "up":
			aim -= c.arg
		}
	}
	return strconv.Itoa(depth * pos)
}

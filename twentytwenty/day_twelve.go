package twentytwenty

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type commandD12 struct {
	command rune
	arg     int
}

func readInputDayTwelve(fp *bufio.Reader) []commandD12 {
	var res []commandD12
	utils.ReadStrings(fp, func(s string) {
		var c rune
		var arg int
		fmt.Sscanf(s, "%c%d", &c, &arg)
		switch c {
		case 'R':
			arg %= 360
			res = append(res, commandD12{command: c, arg: arg / 90})
		case 'L':
			arg %= 360
			arg = 360 - arg
			res = append(res, commandD12{command: 'R', arg: arg / 90})
		default:
			res = append(res, commandD12{command: c, arg: arg})
		}
	})
	return res
}

func DayTwelveA(fp *bufio.Reader) string {
	commands := readInputDayTwelve(fp)

	// N = 0, E = 1, S = 2, W = 3
	heading := 1
	start := types.Coord{}
	pos := start
	for _, command := range commands {
		switch command.command {
		case 'N':
			pos = pos.UpBy(command.arg)
		case 'S':
			pos = pos.DownBy(command.arg)
		case 'E':
			pos = pos.RightBy(command.arg)
		case 'W':
			pos = pos.LeftBy(command.arg)
		case 'F':
			switch heading {
			case 0:
				pos = pos.UpBy(command.arg)
			case 1:
				pos = pos.RightBy(command.arg)
			case 2:
				pos = pos.DownBy(command.arg)
			case 3:
				pos = pos.LeftBy(command.arg)
			}
		case 'R':
			heading = (heading + command.arg) % 4
		}
	}
	return strconv.Itoa(utils.ManhattanDistance(start, pos))
}

func DayTwelveB(fp *bufio.Reader) string {
	commands := readInputDayTwelve(fp)

	start := types.Coord{}
	pos := start
	wpxOffset := 10
	wpyOffset := 1
	for _, command := range commands {
		switch command.command {
		case 'N':
			wpyOffset += command.arg
		case 'S':
			wpyOffset -= command.arg
		case 'E':
			wpxOffset += command.arg
		case 'W':
			wpxOffset -= command.arg
		case 'F':
			pos = types.Coord{
				X: pos.X + (wpxOffset * command.arg),
				Y: pos.Y + (wpyOffset * command.arg)}
		case 'R':
			for i := 0; i < command.arg; i++ {
				wpxOffset, wpyOffset = wpyOffset, wpxOffset
				wpyOffset *= -1
			}
		}
	}
	return strconv.Itoa(utils.ManhattanDistance(start, pos))
}

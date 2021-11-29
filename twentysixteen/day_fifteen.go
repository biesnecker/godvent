package twentysixteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type diskD15 struct {
	id, positions, starting int
}

func readInputDayFifteen(fp *bufio.Reader) []diskD15 {
	var res []diskD15
	utils.ReadStrings(fp, func(s string) {
		// Disc #1 has 17 positions; at time=0, it is at position 1.
		var id, positions, starting int
		fmt.Sscanf(
			s,
			"Disc #%d has %d positions; at time=0, it is at position %d.",
			&id,
			&positions,
			&starting)
		res = append(res, diskD15{id: id, positions: positions, starting: starting})
	})
	return res
}

func findSolutionDayFifteen(disks []diskD15) int {
	time := disks[0].positions - ((disks[0].id + disks[0].starting) % disks[0].positions)
	step := disks[0].positions
outerloop:
	for {
		for i := range disks {
			if (time+disks[i].id+disks[i].starting)%disks[i].positions != 0 {
				time += step
				continue outerloop
			}
		}
		break
	}
	return time
}

func DayFifteenA(fp *bufio.Reader) string {
	disks := readInputDayFifteen(fp)
	return strconv.Itoa(findSolutionDayFifteen(disks))
}

func DayFifteenB(fp *bufio.Reader) string {
	disks := readInputDayFifteen(fp)
	nextId := len(disks) + 1
	disks = append(disks, diskD15{id: nextId, positions: 11, starting: 0})
	return strconv.Itoa(findSolutionDayFifteen(disks))
}

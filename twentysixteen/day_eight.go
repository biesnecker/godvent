package twentysixteen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

const SCREEN_WIDTH int = 50
const SCREEN_HEIGHT int = 6

func readInputDayEight(fp *bufio.Reader) [SCREEN_HEIGHT][SCREEN_WIDTH]bool {
	var screen [SCREEN_HEIGHT][SCREEN_WIDTH]bool
	utils.ReadStrings(fp, func(s string) {
		if strings.HasPrefix(s, "rect") {
			var x, y int
			fmt.Sscanf(s, "rect %dx%d", &x, &y)
			for i := 0; i < y; i++ {
				for j := 0; j < x; j++ {
					screen[i][j] = true
				}
			}
		} else {
			var id, amount int
			if strings.Contains(s, "row") {
				// Rotate a row
				fmt.Sscanf(s, "rotate row y=%d by %d", &id, &amount)
				var buffer [SCREEN_WIDTH]bool

				for i, lit := range screen[id] {
					if lit {
						newIdx := (i + amount) % SCREEN_WIDTH
						buffer[newIdx] = true
					}
				}
				// Write it back
				for i := range buffer {
					screen[id][i] = buffer[i]
				}
			} else {
				// Rotate a column
				fmt.Sscanf(s, "rotate column x=%d by %d", &id, &amount)
				var buffer [SCREEN_HEIGHT]bool

				for i, row := range screen {
					newIdx := (i + amount) % SCREEN_HEIGHT
					buffer[newIdx] = row[id]
				}
				// Write it back
				for i := range buffer {
					screen[i][id] = buffer[i]
				}
			}
		}
	})

	return screen
}

func DayEightA(fp *bufio.Reader) string {
	total := 0
	screen := readInputDayEight(fp)
	for i := range screen {
		for j := range screen[i] {
			if screen[i][j] {
				total++
			}
		}
	}
	return strconv.Itoa(total)
}

func DayEightB(fp *bufio.Reader) string {
	screen := readInputDayEight(fp)
	lines := make([]string, 0, SCREEN_HEIGHT)
	for i := range screen {
		v := make([]byte, 0, SCREEN_WIDTH)
		for j := range screen[i] {
			if screen[i][j] {
				v = append(v, '#')
			} else {
				v = append(v, ' ')
			}
		}
		lines = append(lines, string(v))
	}
	return strings.Join(lines, "\n")
}

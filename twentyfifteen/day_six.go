package twentyfifteen

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

const (
	toggleLights  int = 0
	turnLightsOn  int = 1
	turnLightsOff int = 2
)

func buildInputHandlerSix(
	handler func(int, types.Coord, types.Coord),
) func(string) {
	return func(s string) {
		var command, x1, x2, y1, y2 int
		words := strings.Fields(s)
		if words[0] == "toggle" {
			command = toggleLights
			words = words[1:]
		} else {
			if words[1] == "on" {
				command = turnLightsOn
			} else {
				command = turnLightsOff
			}
			words = words[2:]
		}
		fmt.Sscanf(words[0], "%d,%d", &x1, &y1)
		fmt.Sscanf(words[2], "%d,%d", &x2, &y2)

		handler(command, types.Coord{X: x1, Y: y1}, types.Coord{X: x2, Y: y2})
	}
}

func DaySixA(fp *bufio.Reader) string {
	var grid = [1000][1000]bool{}

	utils.ReadStrings(fp, buildInputHandlerSix(
		func(command int, from types.Coord, to types.Coord) {
			for i := from.X; i <= to.X; i++ {
				for j := from.Y; j <= to.Y; j++ {
					switch command {
					case toggleLights:
						grid[i][j] = !grid[i][j]
					case turnLightsOn:
						grid[i][j] = true
					case turnLightsOff:
						grid[i][j] = false
					}
				}
			}
		}))

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				count++
			}
		}
	}
	return strconv.Itoa(count)
}

func DaySixB(fp *bufio.Reader) string {
	var grid = [1000][1000]int{}

	utils.ReadStrings(fp, buildInputHandlerSix(
		func(command int, from types.Coord, to types.Coord) {
			for i := from.X; i <= to.X; i++ {
				for j := from.Y; j <= to.Y; j++ {
					switch command {
					case toggleLights:
						grid[i][j] += 2
					case turnLightsOn:
						grid[i][j] += 1
					case turnLightsOff:
						if grid[i][j] > 0 {
							grid[i][j] -= 1
						}
					}
				}
			}
		}))

	count := 0
	for i := range grid {
		for j := range grid[i] {
			count += grid[i][j]
		}
	}
	return strconv.Itoa(count)
}

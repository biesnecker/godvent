package twentyeighteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type inputD3 struct {
	id   int
	x, y int
	w, h int
}

func readInputDayThree(fp *bufio.Reader) []inputD3 {
	var res []inputD3
	utils.ReadStrings(fp, func(s string) {
		var i inputD3
		fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &i.id, &i.x, &i.y, &i.w, &i.h)
		res = append(res, i)
	})
	return res
}

func buildGridDayThree(input []inputD3) [1200][1200]int {
	var grid [1200][1200]int
	for x := range input {
		in := &input[x]
		for i := in.x; i < in.x+in.w; i++ {
			for j := in.y; j < in.y+in.h; j++ {
				grid[i][j]++
			}
		}
	}
	return grid
}

func DayThreeA(fp *bufio.Reader) string {
	input := readInputDayThree(fp)
	grid := buildGridDayThree(input)

	total := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				total++
			}
		}
	}
	return strconv.Itoa(total)
}

func DayThreeB(fp *bufio.Reader) string {
	input := readInputDayThree(fp)
	grid := buildGridDayThree(input)

nextinput:
	for x := range input {
		in := &input[x]
		found := true

		for i := in.x; i < in.x+in.w; i++ {
			for j := in.y; j < in.y+in.h; j++ {
				if grid[i][j] > 1 {
					found = false
					continue nextinput
				}
			}
		}

		if found {
			return strconv.Itoa(in.id)
		}
	}

	return ""
}

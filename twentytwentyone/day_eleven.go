package twentytwentyone

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/types/queue"
	"github.com/biesnecker/godvent/utils"
)

type octopus struct {
	loc     types.Coord
	energy  int
	flashed bool
}

func readInputDayEleven(fp *bufio.Reader) map[types.Coord]int {
	res := make(map[types.Coord]int)
	utils.ReadStringsWithIndex(fp, func(i int, s string) {
		for j := range s {
			res[types.Coord{X: i, Y: j}] = int(s[j] - '0')
		}
	})
	return res
}

func DayElevenA(fp *bufio.Reader) string {
	input := readInputDayEleven(fp)
	flashes := 0
	for step := 0; step < 100; step++ {
		q := queue.New()
		for loc, o := range input {
			input[loc]++
			if o == 9 {
				q.Push(loc)
			}
		}
		for !q.Empty() {
			loc := q.Pop().(types.Coord)
			e := input[loc]
			if e == 0 {
				continue
			}
			for _, otherloc := range loc.GetSurroundingCoords() {
				if other, ok := input[otherloc]; ok {
					if other == 0 {
						continue
					}
					input[otherloc]++
					if other == 9 {
						q.Push(otherloc)
					}
				}
			}
			input[loc] = 0
			flashes++
		}

	}
	return strconv.Itoa(flashes)
}

func DayElevenB(fp *bufio.Reader) string {
	input := readInputDayEleven(fp)
	flashes := 0
	step := 0
	for {
		step++
		flashes = 0
		q := queue.New()
		for loc, o := range input {
			if o == 9 {
				q.Push(loc)
			}
			input[loc]++
		}
		for !q.Empty() {
			loc := q.Pop().(types.Coord)
			e := input[loc]
			if e == 0 {
				continue
			}
			for _, otherloc := range loc.GetSurroundingCoords() {
				if other, ok := input[otherloc]; ok {
					if other == 0 {
						continue
					}
					if other == 9 {
						q.Push(otherloc)
					}
					input[otherloc]++
				}
			}
			input[loc] = 0
			flashes++
		}

		if flashes == len(input) {
			break
		}
	}
	return strconv.Itoa(step)
}

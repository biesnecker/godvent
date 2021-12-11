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

func readInputDayEleven(fp *bufio.Reader) map[types.Coord]*octopus {
	res := make(map[types.Coord]*octopus)
	utils.ReadStringsWithIndex(fp, func(i int, s string) {
		for j := range s {
			e := int(s[j] - '0')
			loc := types.Coord{X: i, Y: j}
			res[loc] = &octopus{energy: e, loc: loc}
		}
	})
	return res
}

func DayElevenA(fp *bufio.Reader) string {
	input := readInputDayEleven(fp)
	flashes := 0
	for step := 0; step < 100; step++ {
		q := queue.New()
		for _, o := range input {
			o.energy++
			if o.energy > 9 {
				q.Push(o)
			}
		}
		for !q.Empty() {
			o := q.Pop().(*octopus)
			if o.flashed {
				continue
			}
			for _, otherloc := range o.loc.GetSurroundingCoords() {
				if other, ok := input[otherloc]; ok {
					if other.flashed {
						continue
					}
					other.energy++
					if other.energy > 9 {
						q.Push(other)
					}
				}
			}
			o.flashed = true
			o.energy = 0
			flashes++
		}

		// Clear the flashed state
		for _, o := range input {
			o.flashed = false
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
		for _, o := range input {
			o.energy++
			if o.energy > 9 {
				q.Push(o)
			}
		}
		for !q.Empty() {
			o := q.Pop().(*octopus)
			if o.flashed {
				continue
			}
			for _, otherloc := range o.loc.GetSurroundingCoords() {
				if other, ok := input[otherloc]; ok {
					if other.flashed {
						continue
					}
					other.energy++
					if other.energy > 9 {
						q.Push(other)
					}
				}
			}
			o.flashed = true
			o.energy = 0
			flashes++
		}

		// Clear the flashed state
		for _, o := range input {
			o.flashed = false
		}
		if flashes == len(input) {
			break
		}
	}
	return strconv.Itoa(step)
}

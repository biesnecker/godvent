package twentytwentyone

import (
	"bufio"
	"math"
	"sort"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/types/queue"
	"github.com/biesnecker/godvent/utils"
)

func readInputDayNine(fp *bufio.Reader) map[types.Coord]int {
	res := make(map[types.Coord]int)
	utils.ReadStringsWithIndex(fp, func(i int, s string) {
		for j := range s {
			res[types.Coord{X: i, Y: j}] = int(s[j] - '0')
		}
	})
	return res
}

func findLowPoints(m map[types.Coord]int) []types.Coord {
	var res []types.Coord
	for c, v := range m {
		lowest := math.MaxInt
		for _, adj := range c.GetAdjacentCoords() {
			if depth, ok := m[adj]; ok {
				if lowest > depth {
					lowest = depth
				}
			}
		}
		if v < lowest {
			res = append(res, c)
		}
	}
	return res
}

func DayNineA(fp *bufio.Reader) string {
	input := readInputDayNine(fp)
	risk := 0
	for _, loc := range findLowPoints(input) {
		risk += input[loc] + 1
	}
	return strconv.Itoa(risk)
}

func DayNineB(fp *bufio.Reader) string {
	input := readInputDayNine(fp)
	claimed := make(map[types.Coord]bool)
	var basins []int
	for _, loc := range findLowPoints(input) {
		q := queue.New()
		q.Push(loc)
		claimed[loc] = true
		size := 1

		for !q.Empty() {
			c := q.Pop().(types.Coord)
			for _, adj := range c.GetAdjacentCoords() {
				if claimed[adj] {
					continue
				}
				if v, ok := input[adj]; !ok {
					continue
				} else if v == 9 {
					continue
				}
				claimed[adj] = true
				q.Push(adj)
				size++
			}
		}

		basins = append(basins, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	prod := 1
	for i := 0; i < 3; i++ {
		prod *= basins[i]
	}
	return strconv.Itoa(prod)
}

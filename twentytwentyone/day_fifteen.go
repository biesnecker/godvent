package twentytwentyone

import (
	"bufio"
	"container/heap"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

func readInputDayFifteen(fp *bufio.Reader) (int, int, map[types.Coord]int) {
	res := make(map[types.Coord]int)
	var maxx, maxy int
	utils.ReadStringsWithIndex(fp, func(y int, s string) {
		if y > maxy {
			maxy = y
		}
		for x := range s {
			if x > maxx {
				maxx = x
			}
			res[types.Coord{X: x, Y: y}] = int(s[x] - '0')
		}
	})
	return maxx, maxy, res

}

type stepD15 struct {
	loc         types.Coord
	risk, index int
}

type priorityQueueD15 struct {
	target types.Coord
	steps  []*stepD15
}

func (pq *priorityQueueD15) Len() int { return len(pq.steps) }

func (pq *priorityQueueD15) Less(i, j int) bool {
	if pq.steps[i].risk == pq.steps[j].risk {
		return utils.ManhattanDistance(pq.target, pq.steps[i].loc) <
			utils.ManhattanDistance(pq.target, pq.steps[j].loc)
	} else {
		return pq.steps[i].risk < pq.steps[j].risk
	}
}

func (pq *priorityQueueD15) Swap(i, j int) {
	pq.steps[i], pq.steps[j] = pq.steps[j], pq.steps[i]
	pq.steps[i].index = i
	pq.steps[j].index = j
}

func (pq *priorityQueueD15) Push(x interface{}) {
	n := len(pq.steps)
	item := x.(*stepD15)
	item.index = n
	pq.steps = append(pq.steps, item)
}

func (pq *priorityQueueD15) Pop() interface{} {
	n := len(pq.steps)
	item := pq.steps[n-1]
	pq.steps[n-1] = nil // avoid memory leak
	item.index = -1     // for safety
	pq.steps = pq.steps[:n-1]
	return item
}

func DayFifteenA(fp *bufio.Reader) string {
	maxx, maxy, input := readInputDayFifteen(fp)

	target := types.Coord{X: maxx, Y: maxy}

	pq := &priorityQueueD15{target: target}
	pq.Push(&stepD15{loc: types.Coord{X: 0, Y: 0}, risk: 0})

	minrisk := math.MaxInt

	risks := make(map[types.Coord]int)

	for pq.Len() != 0 {
		s := heap.Pop(pq).(*stepD15)

		if s.loc == target {
			if minrisk > s.risk {
				minrisk = s.risk
				continue
			}
		}

	adjloop:
		for _, nextloc := range s.loc.GetAdjacentCoords() {
			if nextrisk, ok := input[nextloc]; ok {
				nextrisk += s.risk
				if nextrisk >= minrisk {
					continue adjloop
				}
				if c, ok := risks[nextloc]; ok && c <= nextrisk {
					continue adjloop
				}
				risks[nextloc] = nextrisk
				heap.Push(pq, &stepD15{loc: nextloc, risk: nextrisk})
			}
		}

	}
	return strconv.Itoa(minrisk)
}

func DayFifteenB(fp *bufio.Reader) string {
	maxx, maxy, input := readInputDayFifteen(fp)

	maxxe := (maxx+1)*5 - 1
	maxye := (maxy+1)*5 - 1

	target := types.Coord{X: maxxe, Y: maxye}

	pq := &priorityQueueD15{target: target}
	pq.Push(&stepD15{loc: types.Coord{X: 0, Y: 0}, risk: 0})

	minrisk := math.MaxInt

	risks := make(map[types.Coord]int)

	for pq.Len() > 0 {
		s := heap.Pop(pq).(*stepD15)

		if s.loc == target {
			if minrisk > s.risk {
				minrisk = s.risk
				continue
			}
		}

	adjloop:
		for _, actualnextloc := range s.loc.GetAdjacentCoords() {
			if actualnextloc.X < 0 || actualnextloc.X > maxxe ||
				actualnextloc.Y < 0 || actualnextloc.Y > maxye {
				continue
			}
			nx := actualnextloc.X % (maxx + 1)
			ny := actualnextloc.Y % (maxy + 1)
			nextloc := types.Coord{X: nx, Y: ny}
			xadd := actualnextloc.X / (maxx + 1)
			yadd := actualnextloc.Y / (maxy + 1)
			add := xadd + yadd
			if baserisk, ok := input[nextloc]; ok {
				nextrisk := baserisk + add
				if nextrisk > 9 {
					nextrisk = nextrisk%10 + 1
				}
				nextrisk += s.risk
				if c, ok := risks[actualnextloc]; ok && c <= nextrisk {
					continue adjloop
				}
				risks[actualnextloc] = nextrisk
				heap.Push(pq, &stepD15{loc: actualnextloc, risk: nextrisk})
			}
		}

	}
	return strconv.Itoa(minrisk)
}

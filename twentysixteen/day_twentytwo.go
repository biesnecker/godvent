package twentysixteen

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/types/queue"
	"github.com/biesnecker/godvent/types/set"
	"github.com/biesnecker/godvent/utils"
	"github.com/biesnecker/godvent/utils/search"
)

type nodeD22 struct {
	loc                        types.Coord
	size, used, available, pct int
}

func readInputDayTwentyTwo(fp *bufio.Reader) []nodeD22 {
	var res []nodeD22
	utils.ReadStringsWithIndex(fp, func(idx int, s string) {
		if idx < 2 {
			return
		}
		fields := strings.Fields(s)
		var x, y int
		fmt.Sscanf(fields[0], "/dev/grid/node-x%d-y%d", &x, &y)
		size := utils.ReadInt(fields[1][:len(fields[1])-1])
		used := utils.ReadInt(fields[2][:len(fields[2])-1])
		available := utils.ReadInt(fields[3][:len(fields[3])-1])
		pct := utils.ReadInt(fields[4][:len(fields[4])-1])

		res = append(res, nodeD22{
			loc:       types.Coord{X: x, Y: y},
			size:      size,
			used:      used,
			available: available,
			pct:       pct})

	})
	return res
}

func getValidPairs(nodes []nodeD22) []nodeD22 {
	var pairs []nodeD22
	var zeroIdx int
	for i := range nodes {
		for j := range nodes {
			if i != j && nodes[i].used > 0 && nodes[i].used <= nodes[j].available {
				pairs = append(pairs, nodes[i])
				zeroIdx = j
			}
		}
	}
	return append(pairs, nodes[zeroIdx])
}

func getNodeMap(nodes []nodeD22) map[types.Coord]nodeD22 {
	m := make(map[types.Coord]nodeD22)
	for i := range nodes {
		m[nodes[i].loc] = nodes[i]
	}
	return m
}

// Uses a BFS to find the shortest path between the current open space (start)
// and the target position into which to move the payload (target), while
// skipping the current position of the payload (avoid). Used to figure out how
// many steps are required to reposition the open space to the next spot into
// which they payload must move.
func getStepsToTarget(
	grid map[types.Coord]nodeD22,
	start types.Coord,
	target types.Coord,
	avoiding types.Coord,
) int {
	type st struct {
		loc   types.Coord
		steps int
	}

	locationQ := queue.New()
	visited := set.New()

	locationQ.Push(st{loc: start, steps: 0})

	for !locationQ.Empty() {
		current := locationQ.Pop().(st)

		if current.loc == target {
			return current.steps
		}

		for _, newLoc := range current.loc.GetAdjacentCoords() {
			if newLoc.IsNegative() || newLoc == avoiding {
				continue
			}
			if _, ok := grid[newLoc]; !ok {
				continue
			}
			if !visited.Contains(newLoc) {
				visited.Insert(newLoc)
				locationQ.Push(st{loc: newLoc, steps: current.steps + 1})
			}
		}
	}
	return -1
}

type stateD22 struct {
	payloadLoc, openLoc types.Coord
	steps               int
}

func (s *stateD22) GetRepr(interface{}) interface{} {
	return [4]int{s.openLoc.X, s.openLoc.Y, s.payloadLoc.X, s.payloadLoc.Y}
}

func (s *stateD22) GetNext(g interface{}) []search.Searchable {
	grid := g.(map[types.Coord]nodeD22)

	res := make([]search.Searchable, 0, 4)

	for _, nextPayloadLoc := range s.payloadLoc.GetAdjacentCoords() {
		if nextPayloadLoc.IsNegative() {
			continue
		}
		nextSteps := 0
		nextOpenLocation := s.payloadLoc
		if _, ok := grid[nextPayloadLoc]; ok {
			if nextPayloadLoc == s.openLoc {
				// We want to move to the currently open space, so it's only
				// one step.
				nextSteps = 1
			} else {
				// Calculate how many steps it will take to get the empty space
				// to the next position, then add one for the actual swap.
				nextSteps = getStepsToTarget(
					grid, s.openLoc, nextPayloadLoc, s.payloadLoc) + 1
			}
			if nextSteps > 0 {
				// It's possible to get there.
				res = append(
					res,
					&stateD22{
						payloadLoc: nextPayloadLoc,
						openLoc:    nextOpenLocation,
						steps:      s.steps + nextSteps})
			}
		}
	}
	return res
}

func DayTwentyTwoA(fp *bufio.Reader) string {
	nodes := readInputDayTwentyTwo(fp)
	return strconv.Itoa(len(getValidPairs(nodes)) - 1)
}

func DayTwentyTwoB(fp *bufio.Reader) string {
	nodes := readInputDayTwentyTwo(fp)
	nodemap := getNodeMap(getValidPairs(nodes))

	var openLoc types.Coord
	var payloadLoc types.Coord

	maxX := math.MinInt64

	for _, v := range nodemap {
		if v.loc.Y == 0 && v.loc.X > maxX {
			maxX = v.loc.X
			payloadLoc = v.loc
		} else if v.used == 0 {
			openLoc = v.loc
		}
	}

	bfs := search.NewBFSGenerator(
		&stateD22{
			payloadLoc: payloadLoc,
			openLoc:    openLoc,
			steps:      0},
		nodemap)

	for p := bfs.Next(); p != nil; p = bfs.Next() {
		p := p.(*stateD22)
		if p.payloadLoc.Equals(0, 0) {
			return strconv.Itoa(p.steps)
		}
	}

	return ""
}

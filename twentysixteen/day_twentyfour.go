package twentysixteen

import (
	"bufio"
	"math"
	"strconv"
	"sync"
	"unicode"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
	"github.com/biesnecker/godvent/utils/search"
)

type gridD24 map[types.Coord]bool
type locsD24 map[int]types.Coord

type locPair struct{ a, b int }
type distancesD24 map[locPair]int

type bfsStateD24 struct {
	grid      gridD24
	locs      locsD24
	distances distancesD24
	wg        sync.WaitGroup
	lock      sync.Mutex
}

type bfsPositionD24 struct {
	loc   types.Coord
	steps int
}

func (p *bfsPositionD24) GetNext(userData interface{}) []search.Searchable {
	grid := userData.(gridD24)
	positions := make([]search.Searchable, 0, 4)
	for _, loc := range p.loc.GetAdjacentCoords() {
		if isOpen, found := grid[loc]; isOpen && found {
			positions = append(positions, &bfsPositionD24{
				loc:   loc,
				steps: p.steps + 1})
		}
	}
	return positions
}

func (p *bfsPositionD24) GetRepr(interface{}) interface{} {
	return p.loc
}

func readInputDayTwentyFour(fp *bufio.Reader) (gridD24, locsD24) {
	grid := make(map[types.Coord]bool)
	locs := make(map[int]types.Coord)

	utils.ReadStringsWithIndex(fp, func(y int, s string) {
		for x, c := range s {
			coord := types.Coord{X: x, Y: y}
			if unicode.IsDigit(rune(c)) {
				startLocation := int(c - '0')
				grid[coord] = true
				locs[startLocation] = coord
			} else if c == '.' {
				grid[coord] = true
			} else {
				grid[coord] = false
			}
		}
	})

	return grid, locs
}

func findPathBetween(state *bfsStateD24, start, end int) {
	state.wg.Add(1)
	defer state.wg.Done()
	startCoord := state.locs[start]
	endCoord := state.locs[end]

	bfs := search.NewBFSGenerator(
		&bfsPositionD24{loc: startCoord, steps: 0},
		state.grid)

	for p := bfs.Next(); p != nil; p = bfs.Next() {
		pos := p.(*bfsPositionD24)
		if pos.loc == endCoord {
			state.lock.Lock()
			state.distances[locPair{start, end}] = pos.steps
			state.distances[locPair{end, start}] = pos.steps
			state.lock.Unlock()
			break
		}
	}
}

func findSolutionDay24(fp *bufio.Reader, returnToStart bool) int {
	grid, locs := readInputDayTwentyFour(fp)
	state := bfsStateD24{
		grid:      grid,
		locs:      locs,
		distances: make(map[locPair]int)}

	for i := 0; i < len(locs)-1; i++ {
		for j := i + 1; j < len(locs); j++ {
			go findPathBetween(&state, i, j)
		}
	}
	state.wg.Wait()

	shortest := math.MaxInt64

	utils.Permutations(len(locs)-1, func(perm []int) bool {
		total := 0

		lastLocation := 0
		for _, location := range perm {
			total += state.distances[locPair{lastLocation, location + 1}]
			lastLocation = location + 1
		}

		if returnToStart {
			total += state.distances[locPair{lastLocation, 0}]
		}

		if total < shortest {
			shortest = total
		}
		return true
	})

	return shortest
}

func DayTwentyFourA(fp *bufio.Reader) string {
	return strconv.Itoa(findSolutionDay24(fp, false))
}

func DayTwentyFourB(fp *bufio.Reader) string {
	return strconv.Itoa(findSolutionDay24(fp, true))
}

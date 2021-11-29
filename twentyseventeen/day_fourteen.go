package twentyseventeen

import (
	"bufio"
	"fmt"
	"math/bits"
	"strconv"

	"github.com/biesnecker/godvent/twentyseventeen/knothash"
	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
	"github.com/biesnecker/godvent/utils/search"
)

func DayFourteenA(fp *bufio.Reader) string {
	salt := utils.ReadSingleString(fp)
	total := 0
	for i := 0; i < 128; i++ {
		input := fmt.Sprintf("%s-%d", salt, i)
		hash := knothash.KnotHash([]byte(input))
		for i := range hash {
			total += bits.OnesCount8(hash[i])
		}
	}
	return strconv.Itoa(total)
}

type gridcellD14 struct {
	loc types.Coord
}

func (g *gridcellD14) GetNext(userData interface{}) []search.Searchable {
	grid := userData.(map[types.Coord]struct{})
	res := make([]search.Searchable, 0, 4)
	for _, dir := range g.loc.GetAdjacentCoords() {
		if _, found := grid[dir]; found {
			res = append(res, &gridcellD14{loc: dir})
		}
	}
	return res
}

func (g *gridcellD14) GetRepr(interface{}) interface{} {
	return g.loc
}

func DayFourteenB(fp *bufio.Reader) string {
	grid := make(map[types.Coord]struct{})
	salt := utils.ReadSingleString(fp)
	for row := 0; row < 128; row++ {
		input := fmt.Sprintf("%s-%d", salt, row)
		hash := knothash.KnotHash([]byte(input))
		for i := range hash {
			for shift := 7; shift >= 0; shift-- {
				if hash[i]&byte(1<<shift) > 0 {
					grid[types.Coord{X: row, Y: i*8 + (7 - shift)}] = struct{}{}
				}
			}
		}
	}
	regions := 0
	for len(grid) > 0 {
		regions++

		var loc types.Coord

		// Get one, doesn't matter which.
		for k := range grid {
			loc = k
			break
		}

		bfs := search.NewBFSGenerator(&gridcellD14{loc}, grid)
		bfs.ForEach(func(elem search.Searchable, userData interface{}) bool {
			delete(grid, elem.(*gridcellD14).loc)
			return true
		})
	}
	return strconv.Itoa(regions)
}

package twentytwentyone

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type sensorField = map[types.Coord3]struct{}

func readInputDay19(fp *bufio.Reader) []sensorField {
	res := make([]sensorField, 0, 48)
	utils.ReadStrings(fp, func(s string) {
		if len(s) == 0 {
			return
		} else if strings.HasPrefix(s, "---") {
			res = append(res, make(sensorField))
		} else {
			var c types.Coord3
			fmt.Sscanf(s, "%d,%d,%d", &c.X, &c.Y, &c.Z)
			res[len(res)-1][c] = struct{}{}
		}
	})
	return res
}

func getRotations() [][]int {
	return [][]int{
		{0, 0, 0}, {90, 0, 0}, {180, 0, 0},
		{270, 0, 0}, {0, 90, 0}, {90, 90, 0},
		{180, 90, 0}, {270, 90, 0}, {0, 180, 0},
		{90, 180, 0}, {180, 180, 0}, {270, 180, 0},
		{0, 270, 0}, {90, 270, 0}, {180, 270, 0},
		{270, 270, 0}, {0, 0, 90}, {90, 0, 90},
		{180, 0, 90}, {270, 0, 90}, {0, 0, 270},
		{90, 0, 270}, {180, 0, 270}, {270, 0, 270},
	}
}

func maybeFindOverlap(known, unknown sensorField) (sensorField, types.Coord3) {
	rotated := make([]types.Coord3, len(unknown))
	translated := make([]types.Coord3, len(unknown))
	for _, r := range getRotations() {
		ui := 0
		for u := range unknown {
			rotated[ui] = u.Rotate(r[0], r[1], r[2])
			ui++
		}
		for k := range known {
			for _, u := range rotated {
				diff := k.Diff(u)
				match := 0
				for i := range rotated {
					translated[i] = rotated[i].Translate(diff)
					if _, ok := known[translated[i]]; ok {
						match++
					}
				}
				if match >= 12 {
					match := make(sensorField)
					for i := range translated {
						match[translated[i]] = struct{}{}
					}
					return match, types.Coord3{}.Translate(diff.Negate())
				}
			}
		}
	}
	return nil, types.Coord3{}
}

func findSolutionD19(fp *bufio.Reader) (sensorField, []types.Coord3) {
	input := readInputDay19(fp)

	scannerLocs := make([]types.Coord3, len(input))

	knownIdxs := []int{0}
	unknownIdxs := make([]int, 0, len(input)-1)
	for i := 1; i < len(input); i++ {
		unknownIdxs = append(unknownIdxs, i)
	}

	for len(unknownIdxs) > 0 {

		var remainingIdxs []int
		var matchedIdxs []int
		var wg sync.WaitGroup
		var matchedLock, remainingLock sync.Mutex
		for _, ui := range unknownIdxs {
			uii := ui
			wg.Add(1)
			go func() {
				defer wg.Done()
				matched := false
				for _, ki := range knownIdxs {
					maybeMatch, scanner := maybeFindOverlap(input[ki], input[uii])
					if maybeMatch != nil {
						input[uii] = maybeMatch
						scannerLocs[uii] = scanner
						matched = true
						break
					}
				}
				if matched {
					matchedLock.Lock()
					matchedIdxs = append(matchedIdxs, uii)
					matchedLock.Unlock()
				} else {
					remainingLock.Lock()
					remainingIdxs = append(remainingIdxs, uii)
					remainingLock.Unlock()
				}
			}()
		}
		wg.Wait()
		if len(remainingIdxs) == len(unknownIdxs) {
			panic("Didn't find any matches")
		}
		knownIdxs = matchedIdxs
		unknownIdxs = remainingIdxs
	}

	unified := make(sensorField)
	for _, f := range input {
		for k := range f {
			unified[k] = struct{}{}
		}
	}

	return unified, scannerLocs
}

func DayNineteenA(fp *bufio.Reader) string {
	beacons, _ := findSolutionD19(fp)
	return strconv.Itoa(len(beacons))
}

func DayNineteenB(fp *bufio.Reader) string {
	_, scanners := findSolutionD19(fp)
	max := 0
	for i := 0; i < len(scanners)-2; i++ {
		for j := i + 1; j < len(scanners)-1; j++ {
			d := utils.ManhattanDistance3(scanners[i], scanners[j])
			if d > max {
				max = d
			}
		}
	}
	return strconv.Itoa(max)
}

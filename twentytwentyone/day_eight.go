package twentytwentyone

import (
	"bufio"
	"fmt"
	"math/bits"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type inputD8 struct {
	left  [10]uint8
	right [4]uint8
}

func maskFromStringRepr(rep string) uint8 {
	mask := uint8(0)
	for _, r := range rep {
		b := byte(r) - 'a'
		mask |= (uint8(1) << b)
	}
	return mask
}

func readInputDayEight(fp *bufio.Reader) []inputD8 {
	var res []inputD8
	utils.ReadStrings(fp, func(s string) {
		var i inputD8
		var left [10]string
		var right [4]string
		fmt.Sscanf(s, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s",
			&left[0], &left[1], &left[2], &left[3], &left[4], &left[5],
			&left[6], &left[7], &left[8], &left[9], &right[0], &right[1],
			&right[2], &right[3],
		)
		for idx, v := range left {
			i.left[idx] = maskFromStringRepr(v)
		}
		for idx, v := range right {
			i.right[idx] = maskFromStringRepr(v)
		}
		res = append(res, i)
	})
	return res
}

func DayEightA(fp *bufio.Reader) string {
	input := readInputDayEight(fp)
	count := 0
	for i := range input {
		for _, v := range input[i].right {
			switch bits.OnesCount8(v) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func DayEightB(fp *bufio.Reader) string {
	var total int
	input := readInputDayEight(fp)

	for _, in := range input {

		segments := make([]uint8, 7)
		resolved := make([]bool, 7)
		for i := 0; i < 7; i++ {
			segments[i] = 0x7f
			resolved[i] = false
		}

		numberReps := make(map[uint8]int)
		for i := range in.left {
			numberReps[in.left[i]] = -1
		}

		unresolved := make([]uint8, len(in.left))
		for i := range in.left {
			unresolved[i] = in.left[i]
		}

		for len(unresolved) > 0 {
			stillUnresolved := make([]uint8, 0, len(unresolved))

			// Update the resolved map.
			for i := 0; i < 7; i++ {
				resolved[i] = bits.OnesCount8(segments[i]) == 1
			}

			for _, liv := range unresolved {
				switch bits.OnesCount8(liv) {
				case 2:
					// Has to be 1
					segments[0] &= ^liv
					segments[1] &= ^liv
					segments[2] &= liv
					segments[3] &= ^liv
					segments[4] &= ^liv
					segments[5] &= liv
					segments[6] &= ^liv
					numberReps[liv] = 1
				case 3:
					// Has to be 7
					segments[0] &= liv
					segments[1] &= ^liv
					segments[2] &= liv
					segments[3] &= ^liv
					segments[4] &= ^liv
					segments[5] &= liv
					segments[6] &= ^liv
					numberReps[liv] = 7
				case 4:
					// Has to be 4
					segments[0] &= ^liv
					segments[1] &= liv
					segments[2] &= liv
					segments[3] &= liv
					segments[4] &= ^liv
					segments[5] &= liv
					segments[6] &= ^liv
					numberReps[liv] = 4
				case 5:
					// Could be 2, 3, or 5.
					// Narrow common fields
					segments[0] &= liv
					segments[3] &= liv
					segments[6] &= liv

					// Check for exact matches
					if resolved[1] && resolved[5] && ((segments[1]|segments[5])&liv) == 0 {
						numberReps[liv] = 2
					} else if resolved[1] && resolved[4] && ((segments[1]|segments[4])&liv) == 0 {
						numberReps[liv] = 3
					} else if resolved[2] && resolved[4] && ((segments[2]|segments[4])&liv) == 0 {
						numberReps[liv] = 5
					} else if resolved[4] && segments[4]&liv != 0 {
						// Only 2 uses segment 4
						numberReps[liv] = 2
						segments[1] &= ^liv
						segments[2] &= liv
						segments[5] &= ^liv
					} else if resolved[1] && segments[1]&liv != 0 {
						// Only 5 uses segment 1
						numberReps[liv] = 5
						segments[2] &= ^liv
						segments[4] &= ^liv
						segments[5] &= liv
					} else {
						stillUnresolved = append(stillUnresolved, liv)
					}
				case 6:
					// Could be 0, 6, or 9
					// Narrow common fields
					segments[0] &= liv
					segments[1] &= liv
					segments[5] &= liv
					segments[6] &= liv

					// We can only solve these if the segment that they're
					// missing is resolved.
					if resolved[3] && segments[3]&liv == 0 {
						numberReps[liv] = 0
					} else if resolved[2] && segments[2]&liv == 0 {
						numberReps[liv] = 6
					} else if resolved[4] && segments[4]&liv == 0 {
						numberReps[liv] = 9
					} else {
						stillUnresolved = append(stillUnresolved, liv)
					}
				case 7:
					numberReps[liv] = 8
				}

			}

			unresolved = stillUnresolved
		}

		// Do the decode.
		n := 0
		for i := range in.right {
			n *= 10
			n += numberReps[in.right[i]]
		}
		total += n

	}
	return strconv.Itoa(total)
}

package twentytwentyone

import (
	"bufio"
	"fmt"
	"math/bits"
	"sort"
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

func sortStringByCharacter(rep string) string {
	b := []byte(rep)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
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
		fmt.Println(i)
	})
	return res
}

func DayEightA(fp *bufio.Reader) string {
	input := readInputDayEight(fp)
	count := 0
	for i := range input {
		for _, v := range input[i].right {
			switch bits.OnesCount8(v) {
			case 2:
				count++
			case 4:
				count++
			case 3:
				count++
			case 7:
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

func maskFromStringRep(rep string) uint8 {
	mask := uint8(0)
	for _, r := range rep {
		b := byte(r) - 'a'
		mask |= (uint8(1) << b)
	}
	return mask
}

func wireFromSegmentValue(v uint8) byte {
	return 'a' + byte(bits.TrailingZeros8(v))
}

func tryEliminateZero(segments map[int]uint8, mask uint8) bool {
	// No matter what, segments 0, 1, 5, and 6 must be in this mask.
	segments[0] &= mask
	segments[1] &= mask
	segments[5] &= mask
	segments[6] &= mask
	// Can only be zero if the third segment is resolved and the corresponding
	// wire isn't in this mask.
	return bits.OnesCount8(segments[3]) == 1 && segments[3]&mask == 0
}

func eliminateOne(segments map[int]uint8, mask uint8) {
	segments[0] &= ^mask
	segments[1] &= ^mask
	segments[2] &= mask
	segments[3] &= ^mask
	segments[4] &= ^mask
	segments[5] &= mask
	segments[6] &= ^mask
}

func tryEliminateTwo(segments map[int]uint8, mask uint8) bool {
	// No matter what, segments 0, 3, and 6 must be in this mask.
	segments[0] &= mask
	segments[3] &= mask
	segments[6] &= mask
	if bits.OnesCount8(segments[1]) == 1 &&
		bits.OnesCount8(segments[5]) == 1 &&
		(segments[1]|segments[5])&mask == 0 {
		// This is definitely two.
		return true
	} else if bits.OnesCount8(segments[4]) == 1 && segments[4]|mask > 0 {
		// if segment 4 is resolved then it has to be two because no other
		// five segment number uses it.
		segments[1] &= ^mask
		segments[2] &= mask
		segments[5] &= ^mask
		return true
	}

	return false
}

func tryEliminateThree(segments map[int]uint8, mask uint8) bool {
	// No matter what, segments 0, 3, and 6 must be in this mask.
	segments[0] &= mask
	segments[3] &= mask
	segments[6] &= mask
	if bits.OnesCount8(segments[1]) == 1 &&
		bits.OnesCount8(segments[4]) == 1 &&
		(segments[1]|segments[4])&mask == 0 {
		// This is definitely three.
		return true
	}
	return false
}

func eliminateFour(segments map[int]uint8, mask uint8) {
	segments[0] &= ^mask
	segments[1] &= mask
	segments[2] &= mask
	segments[3] &= mask
	segments[4] &= ^mask
	segments[5] &= mask
	segments[6] &= ^mask
}

func tryEliminateFive(segments map[int]uint8, mask uint8) bool {
	// No matter what, segments 0, 3, and 6 must be in this mask.
	segments[0] &= mask
	segments[3] &= mask
	segments[6] &= mask
	if bits.OnesCount8(segments[2]) == 1 &&
		bits.OnesCount8(segments[4]) == 1 &&
		(segments[2]|segments[4])&mask == 0 {
		// This is definitely five.
		return true
	} else if bits.OnesCount8(segments[1]) == 1 && segments[1]|mask > 0 {
		// if segment 1 is resolved then it has to be five because no other
		// five segment number uses it.
		segments[2] &= ^mask
		segments[4] &= ^mask
		segments[5] &= mask
		return true
	}
	return false
}

func tryEliminateSix(segments map[int]uint8, mask uint8) bool {
	// Can only be six if the second segment is resolved and the corresponding
	// wire isn't in this mask.
	return bits.OnesCount8(segments[2]) == 1 && segments[2]&mask == 0
}

func eliminateSeven(segments map[int]uint8, mask uint8) {
	segments[0] &= mask
	segments[1] &= ^mask
	segments[2] &= mask
	segments[3] &= ^mask
	segments[4] &= ^mask
	segments[5] &= mask
	segments[6] &= ^mask
}

func tryEliminateNine(segments map[int]uint8, mask uint8) bool {
	// Can only be nine if the fourth segment is resolved and the corresponding
	// wire isn't in this mask.
	return bits.OnesCount8(segments[4]) == 1 && segments[4]&mask == 0
}

func DayEightB(fp *bufio.Reader) string {
	var total int
	input := readInputDayEight(fp)

	for lineno, in := range input {

		segments := make(map[int]uint8)
		for i := 0; i < 7; i++ {
			segments[i] = 0x7f
		}

		numberReps := make(map[uint8]int)
		for i := range in.left {
			numberReps[in.left[i]] = -1
		}

		for {
			complete := true
			for _, v := range numberReps {
				if v == -1 {
					complete = false
					break
				}
			}
			if complete {
				// We're done. Should be able to decode.
				break
			}

			for _, liv := range in.left {
				switch bits.OnesCount8(liv) {
				case 2:
					eliminateOne(segments, liv)
					fmt.Println(lineno, ": Eliminated 1 : ", segments, numberReps)
					numberReps[liv] = 1
				case 3:
					eliminateSeven(segments, liv)
					fmt.Println(lineno, ": Eliminated 7 : ", segments, numberReps)
					numberReps[liv] = 7
				case 4:
					eliminateFour(segments, liv)
					fmt.Println(lineno, ": Eliminated 4 : ", segments, numberReps)
					numberReps[liv] = 4
				case 5:
					if tryEliminateTwo(segments, liv) {
						fmt.Println(lineno, ": Eliminated 2 : ", segments, numberReps)
						numberReps[liv] = 2
					} else if tryEliminateThree(segments, liv) {
						fmt.Println(lineno, ": Eliminated 3 : ", segments, numberReps)
						numberReps[liv] = 3
					} else if tryEliminateFive(segments, liv) {
						fmt.Println(lineno, ": Eliminated 5 : ", segments, numberReps)
						numberReps[liv] = 5
					}
				case 6:
					if tryEliminateZero(segments, liv) {
						fmt.Println(lineno, ": Eliminated 0 : ", segments, numberReps)
						numberReps[liv] = 0
					} else if tryEliminateSix(segments, liv) {
						fmt.Println(lineno, ": Eliminated 6 : ", segments, numberReps)
						numberReps[liv] = 6
					} else if tryEliminateNine(segments, liv) {
						fmt.Println(lineno, ": Eliminated 9 : ", segments, numberReps)
						numberReps[liv] = 9
					}
				case 7:
					fmt.Println(lineno, ": Eliminated 8 : ", segments, numberReps)
					numberReps[liv] = 8

				}

			}

			// Do the decode.
			n := 0
			for i := range in.right {
				n *= 10
				n += numberReps[in.right[i]]
			}
			total += n
		}
	}
	return strconv.Itoa(total)
}

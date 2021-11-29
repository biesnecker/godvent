package twentyseventeen

import (
	"bufio"
	"hash/maphash"
	"log"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func findSolutionDaySix(fp *bufio.Reader) (int, int) {
	total := 0
	loopIdx := 0
	utils.ReadDelimitedIntegerStrings(fp, "\t", func(b []int) {
		hash := maphash.Hash{}
		hash.SetSeed(maphash.MakeSeed())

		seen := make(map[uint64]int)

		banks := make([]byte, len(b))
		for i, v := range b {
			banks[i] = byte(v)
		}

		cycles := 0

		for {
			cycles++
			highestValue := byte(0)
			highestValueIndex := 0
			for i, v := range banks {
				if v > highestValue {
					highestValue = v
					highestValueIndex = i
				}
			}
			banks[highestValueIndex] = 0
			nextIdx := highestValueIndex
			for j := 0; j < int(highestValue); j++ {
				nextIdx = (nextIdx + 1) % len(banks)
				banks[nextIdx]++
			}

			if _, err := hash.Write(banks); err != nil {
				log.Fatalln(err)
			}
			hv := hash.Sum64()
			if lpidx, ok := seen[hv]; ok {
				// This is a repeat.
				loopIdx = cycles - lpidx
				total = cycles
				break
			} else {
				seen[hv] = cycles
			}
			hash.Reset()
		}
	})
	return total, loopIdx
}

func DaySixA(fp *bufio.Reader) string {
	total, _ := findSolutionDaySix(fp)
	return strconv.Itoa(total)
}

func DaySixB(fp *bufio.Reader) string {
	_, loopSize := findSolutionDaySix(fp)
	return strconv.Itoa(loopSize)
}

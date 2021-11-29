package twentyseventeen

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/biesnecker/godvent/utils"
)

func readInputDayThirteen(fp *bufio.Reader) map[int]int {
	res := make(map[int]int)
	utils.ReadStrings(fp, func(s string) {
		var slot, depth int
		fmt.Sscanf(s, "%d: %d", &slot, &depth)
		res[slot] = depth
	})
	return res
}

func DayThirteenA(fp *bufio.Reader) string {
	scanners := readInputDayThirteen(fp)
	checked := 0
	severity := 0
	time := 0

	for checked < len(scanners) {
		if depth, found := scanners[time]; found {
			checked++
			if time%((depth-1)*2) == 0 {
				severity += time * depth
			}
		}
		time++
	}
	return strconv.Itoa(severity)
}

func DayThirteenB(fp *bufio.Reader) string {
	scanners := readInputDayThirteen(fp)

	var mtx sync.Mutex
	var wg sync.WaitGroup

	minimum := math.MaxInt64
	chunk := 0
	for {
		for i := 0; i < 10; i++ {
			startIdx := chunk*100000 + i*10000
			endIdx := chunk*100000 + (i+1)*10000
			wg.Add(1)
			go func() {
			candidateLoop:
				for delay := startIdx; delay < endIdx; delay++ {
					time := 0
					checked := 0
					for checked < len(scanners) {
						if depth, found := scanners[time]; found {
							checked++
							if (time+delay)%((depth-1)*2) == 0 {
								continue candidateLoop
							}
						}
						time++
					}
					mtx.Lock()
					if delay < minimum {
						minimum = delay
					}
					mtx.Unlock()
				}
				wg.Done()
			}()
		}
		wg.Wait()
		if minimum != math.MaxInt64 {
			return strconv.Itoa(minimum)
		}
		chunk++
	}
}

package twentynineteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentynineteen/intcode"
	"github.com/biesnecker/godvent/utils"
)

func DaySevenA(fp *bufio.Reader) string {
	prog := utils.ReadDeliminatedInts(fp, ",")
	max := 0
	utils.Permutations(5, func(phases []int) bool {
		lastOut := 0
		for i := range phases {
			readC := make(chan int)
			writeC := make(chan int)
			go intcode.Run(prog, writeC, readC)
			writeC <- phases[i]
			writeC <- lastOut
			lastOut = <-readC
		}
		if lastOut > max {
			max = lastOut
		}
		return true
	})
	return strconv.Itoa(max)
}

func DaySevenB(fp *bufio.Reader) string {
	prog := utils.ReadDeliminatedInts(fp, ",")
	max := 0

	utils.Permutations(5, func(phases []int) bool {
		loopbackC := make(chan int, 2)
		loopbackC <- phases[0] + 5
		loopbackC <- 0
		prevOutC := loopbackC
		for _, phase := range phases[1:] {
			outC := make(chan int, 1)
			outC <- phase + 5
			go intcode.Run(prog, prevOutC, outC)
			prevOutC = outC
		}

		intcode.Run(prog, prevOutC, loopbackC)

		val := <-loopbackC
		if val > max {
			max = val
		}
		return true
	})
	return strconv.Itoa(max)
}

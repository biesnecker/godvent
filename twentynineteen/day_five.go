package twentynineteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentynineteen/intcode"
	"github.com/biesnecker/godvent/utils"
)

func DayFiveA(fp *bufio.Reader) string {
	writeC := make(chan int)
	readC := make(chan int)
	prog := utils.ReadDeliminatedInts(fp, ",")

	go func() {
		for {
			writeC <- 1
		}
	}()
	go intcode.Run(prog, writeC, readC)
	lastValue := -1
	for i := range readC {
		lastValue = i
	}
	return strconv.Itoa(lastValue)
}

func DayFiveB(fp *bufio.Reader) string {
	writeC := make(chan int)
	readC := make(chan int)
	prog := utils.ReadDeliminatedInts(fp, ",")

	go func() {
		for {
			writeC <- 5
		}
	}()
	go intcode.Run(prog, writeC, readC)
	lastValue := -1
	for i := range readC {
		lastValue = i
	}
	return strconv.Itoa(lastValue)
}

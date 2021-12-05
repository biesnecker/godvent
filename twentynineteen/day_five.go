package twentynineteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentynineteen/intcode"
)

func DayFiveA(fp *bufio.Reader) string {
	writeC := make(chan int)
	readC := make(chan int)
	ic := intcode.IntcodeInterpreterFromFile(fp)

	go func() {
		for {
			writeC <- 1
		}
	}()
	go ic.Run(writeC, readC)
	lastValue := -1
	for i := range readC {
		lastValue = i
	}
	return strconv.Itoa(lastValue)
}

func DayFiveB(fp *bufio.Reader) string {
	writeC := make(chan int)
	readC := make(chan int)
	ic := intcode.IntcodeInterpreterFromFile(fp)

	go func() {
		for {
			writeC <- 5
		}
	}()
	go ic.Run(writeC, readC)
	lastValue := -1
	for i := range readC {
		lastValue = i
	}
	return strconv.Itoa(lastValue)
}

package twentysixteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentysixteen/asmbunny"
)

func DayTwentyFiveA(fp *bufio.Reader) string {
	prog := asmbunny.ReadAsmBunnyProgram(fp)
	idx := 0

	for {
		outChannel := make(chan int, 10)
		cpu := asmbunny.AsmBunnyCPU{Out: outChannel}
		go cpu.RunProgram(prog, idx, 0, 0, 0)

		expected := 0
		found := false
		for {
			if actual, ok := <-outChannel; !ok {
				found = true
				break
			} else {
				if actual != expected {
					break
				}
				expected = (expected + 1) % 2
			}
		}
		if found {
			return strconv.Itoa(idx)
		}
		idx++
	}
}

func DayTwentyFiveB(fp *bufio.Reader) string {
	return "There is no 2016 Day 25 Part B puzzle"
}

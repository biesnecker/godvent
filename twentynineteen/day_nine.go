package twentynineteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentynineteen/intcode"
	"github.com/biesnecker/godvent/utils"
)

func DayNineA(fp *bufio.Reader) string {
	writeC := make(chan int)
	readC := make(chan int)
	prog := utils.ReadDeliminatedInts(fp, ",")

	go func() { writeC <- 1 }()
	go intcode.Run(prog, writeC, readC)

	return strconv.Itoa(<-readC)
}

func DayNineB(fp *bufio.Reader) string {
	return ""
}

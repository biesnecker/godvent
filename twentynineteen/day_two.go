package twentynineteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentynineteen/intcode"
	"github.com/biesnecker/godvent/utils"
)

const dayTwoBTarget int = 19690720

func DayTwoA(fp *bufio.Reader) string {
	prog := utils.ReadDeliminatedInts(fp, ",")
	prog[1] = 12
	prog[2] = 2
	mem := intcode.Run(prog, nil, nil)

	return strconv.Itoa(mem[0])
}

func DayTwoB(fp *bufio.Reader) string {
	prog := utils.ReadDeliminatedInts(fp, ",")
	for noun := 0; noun < len(prog); noun++ {
		for verb := 0; verb < len(prog); verb++ {
			prog[1] = noun
			prog[2] = verb
			mem := intcode.Run(prog, nil, nil)
			if mem[0] == dayTwoBTarget {
				return strconv.Itoa(100*noun + verb)
			}
		}
	}
	return ""
}

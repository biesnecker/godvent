package twentynineteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentynineteen/intcode"
)

const dayTwoBTarget int = 19690720

func DayTwoA(fp *bufio.Reader) string {
	ic := intcode.IntcodeInterpreterFromFile(fp)
	ic.Set(1, 12)
	ic.Set(2, 2)
	ic.Run()

	return strconv.Itoa(ic.Prog[0])
}

func DayTwoB(fp *bufio.Reader) string {
	prog := intcode.ReadIntcodeProgramFromFile(fp)
	for noun := 0; noun < len(prog); noun++ {
		for verb := 0; verb < len(prog); verb++ {
			progcopy := make([]int, len(prog))
			copy(progcopy, prog)
			i := intcode.Intcode{Prog: progcopy}
			i.Set(1, noun)
			i.Set(2, verb)
			i.Run()
			if i.Get(0) == dayTwoBTarget {
				return strconv.Itoa(100*noun + verb)
			}
		}
	}
	return ""
}

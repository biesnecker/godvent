package twentysixteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentysixteen/asmbunny"
)

func DayTwentyThreeA(fp *bufio.Reader) string {
	prog := asmbunny.ReadAsmBunnyProgram(fp)
	cpu := asmbunny.AsmBunnyCPU{}
	regs := cpu.RunProgram(prog, 7, 0, 0, 0)
	return strconv.Itoa(regs[0])
}

func DayTwentyThreeB(fp *bufio.Reader) string {
	prog := asmbunny.ReadAsmBunnyProgram(fp)
	cpu := asmbunny.AsmBunnyCPU{}
	regs := cpu.RunProgram(prog, 12, 0, 0, 0)
	return strconv.Itoa(regs[0])
}

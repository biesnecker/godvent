package twentysixteen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentysixteen/asmbunny"
)

func DayTwelveA(fp *bufio.Reader) string {
	prog := asmbunny.ReadAsmBunnyProgram(fp)
	cpu := asmbunny.AsmBunnyCPU{}
	regs := cpu.RunProgram(prog, 0, 0, 0, 0)
	return strconv.Itoa(regs[0])
}

func DayTwelveB(fp *bufio.Reader) string {
	prog := asmbunny.ReadAsmBunnyProgram(fp)
	cpu := asmbunny.AsmBunnyCPU{}
	regs := cpu.RunProgram(prog, 0, 0, 1, 0)
	return strconv.Itoa(regs[0])
}

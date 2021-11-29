package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentyseventeen/cpu"
	"github.com/biesnecker/godvent/types/deque"
)

func DayEighteenA(fp *bufio.Reader) string {
	prog := cpu.ReadInputProgram(fp)
	q := deque.New()
	cpu := cpu.NewCpuWithQueues(prog, q, q)

	solution, _ := cpu.ExecuteProgram(true)
	return strconv.Itoa(solution)
}

func DayEighteenB(fp *bufio.Reader) string {
	prog := cpu.ReadInputProgram(fp)

	outputZero := deque.New()
	outputOne := deque.New()

	cpuZero := cpu.NewCpuWithQueues(prog, outputOne, outputZero)
	cpuZero.SetRegister('p', 0)

	cpuOne := cpu.NewCpuWithQueues(prog, outputZero, outputOne)
	cpuOne.SetRegister('p', 1)

	for {
		_, zeroMadeProgress := cpuZero.ExecuteProgram(false)
		_, oneMadeProgress := cpuOne.ExecuteProgram(false)

		if !zeroMadeProgress && !oneMadeProgress {
			break
		}

	}
	return strconv.Itoa(cpuOne.GetSendCount())
}

package intcode

import (
	"bufio"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type Intcode struct {
	Prog []int
}

func ReadIntcodeProgramFromFile(fp *bufio.Reader) []int {
	var prog []int
	in := utils.ReadSingleString(fp)
	for _, i := range strings.Split(in, ",") {
		prog = append(prog, utils.ReadInt(i))
	}
	return prog
}

func IntcodeInterpreterFromFile(fp *bufio.Reader) Intcode {
	return Intcode{Prog: ReadIntcodeProgramFromFile(fp)}
}

func (interp *Intcode) Set(idx, value int) {
	interp.Prog[idx] = value
}

func (interp *Intcode) Get(idx int) int {
	return interp.Prog[idx]
}

func (interp *Intcode) getIndirect(idx int) int {
	i := interp.Prog[idx]
	return interp.Prog[i]
}

func (interp *Intcode) getDest(idx int) *int {
	i := interp.Prog[idx]
	return &interp.Prog[i]
}

func (interp *Intcode) Run() {
	pc := 0
dispatch:
	for {
		switch interp.Prog[pc] {
		case 1:
			opA := interp.getIndirect(pc + 1)
			opB := interp.getIndirect(pc + 2)
			*interp.getDest(pc + 3) = opA + opB
			pc += 4
		case 2:
			opA := interp.getIndirect(pc + 1)
			opB := interp.getIndirect(pc + 2)
			*interp.getDest(pc + 3) = opA * opB
			pc += 4
		case 99:
			break dispatch
		}
	}
}

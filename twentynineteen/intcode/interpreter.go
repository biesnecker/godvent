package intcode

import (
	"bufio"

	"github.com/biesnecker/godvent/utils"
)

type AddressMode = int

const (
	PositionMode  AddressMode = 0
	ImmediateMode AddressMode = 1
)

type DecodedInstruction struct {
	op    int
	modes [3]AddressMode
}

func decodeInstruction(i int) DecodedInstruction {
	d := DecodedInstruction{}
	d.op = i % 100
	i /= 100
	d.modes[0] = i % 10
	i /= 10
	d.modes[1] = i % 10
	i /= 10
	d.modes[2] = i
	return d
}

type Intcode struct {
	Prog []int
}

func IntcodeInterpreterFromFile(fp *bufio.Reader) Intcode {
	return Intcode{Prog: utils.ReadDeliminatedInts(fp, ",")}
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

func (interp *Intcode) getWithMode(idx int, mode AddressMode) int {
	if mode == ImmediateMode {
		return interp.Prog[idx]
	} else {
		return interp.getIndirect(idx)
	}
}

func (interp *Intcode) getDest(idx int) *int {
	i := interp.Prog[idx]
	return &interp.Prog[i]
}

func (interp *Intcode) Run(inC <-chan int, outC chan<- int) {
	pc := 0
dispatch:
	for {
		i := decodeInstruction(interp.Prog[pc])
		switch i.op {
		case 1:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			*interp.getDest(pc + 3) = opA + opB
			pc += 4
		case 2:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			*interp.getDest(pc + 3) = opA * opB
			pc += 4
		case 3:
			input := <-inC
			*interp.getDest(pc + 1) = input
			pc += 2
		case 4:
			outC <- interp.getWithMode(pc+1, i.modes[0])
			pc += 2
		case 5:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			if opA != 0 {
				pc = opB
			} else {
				pc += 3
			}
		case 6:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			if opA == 0 {
				pc = opB
			} else {
				pc += 3
			}
		case 7:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			d := interp.getDest(pc + 3)
			if opA < opB {
				*d = 1
			} else {
				*d = 0
			}
			pc += 4
		case 8:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			d := interp.getDest(pc + 3)
			if opA == opB {
				*d = 1
			} else {
				*d = 0
			}
			pc += 4
		case 99:
			break dispatch
		}
	}
	close(outC)
}

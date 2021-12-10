package intcode

import (
	"log"
)

type AddressMode = int

const (
	PositionMode  AddressMode = 0
	ImmediateMode AddressMode = 1
	RelativeMode  AddressMode = 2
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

type intcode struct {
	mem    map[int]int
	base   int
	input  <-chan int
	output chan<- int
}

func (interp *intcode) getIndirect(idx, offset int) int {
	i := interp.mem[idx+offset]
	return interp.mem[i]
}

func (interp *intcode) getWithMode(idx int, mode AddressMode) int {
	switch mode {
	case PositionMode:
		return interp.getIndirect(idx, 0)
	case ImmediateMode:
		return interp.mem[idx]
	case RelativeMode:
		return interp.getIndirect(idx, interp.base)
	default:
		log.Fatalln("Unknown get mode: ", mode)
	}
	return 0
}

func (interp *intcode) setWithMode(idx, value int, mode AddressMode) {
	var i int
	switch mode {
	case PositionMode:
		i = idx
	case RelativeMode:
		i = idx + interp.base
	default:
		log.Fatalln("Unknown set mode: ", mode)
	}
	interp.mem[i] = value
}

func Run(prog []int, inC <-chan int, outC chan<- int) map[int]int {
	mem := make(map[int]int)
	for i, instr := range prog {
		mem[i] = instr
	}
	interp := intcode{
		mem:    mem,
		base:   0,
		input:  inC,
		output: outC,
	}
	pc := 0
dispatch:
	for {
		i := decodeInstruction(interp.mem[pc])
		switch i.op {
		case 1:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			interp.setWithMode(pc+3, opA+opB, i.modes[2])
			pc += 4
		case 2:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			interp.setWithMode(pc+3, opA*opB, i.modes[2])
			pc += 4
		case 3:
			interp.setWithMode(pc+1, <-inC, i.modes[0])
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
			if opA < opB {
				interp.setWithMode(pc+3, 1, i.modes[2])
			} else {
				interp.setWithMode(pc+3, 0, i.modes[2])
			}
			pc += 4
		case 8:
			opA := interp.getWithMode(pc+1, i.modes[0])
			opB := interp.getWithMode(pc+2, i.modes[1])
			if opA == opB {
				interp.setWithMode(pc+3, 1, i.modes[2])
			} else {
				interp.setWithMode(pc+3, 0, i.modes[2])
			}
			pc += 4
		case 9:
			opA := interp.getWithMode(pc+1, i.modes[0])
			interp.base += opA
			pc += 2
		case 99:
			break dispatch
		}
	}
	if outC != nil {
		close(outC)
	}

	return interp.mem
}

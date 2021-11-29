package asmbunny

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/utils"
)

type AsmBunnyOp int

const (
	CPY AsmBunnyOp = iota
	JNZ
	INC
	DEC
	TGL
	OUT
)

type AsmBunnyCPU struct {
	registers []int
	pc        int
	prog      []AsmBunnyInstruction
	outCnt    int
	Out       chan<- int
}

func (cpu *AsmBunnyCPU) resolveValue(arg AsmBunnyArgument) int {
	if arg.isRegister() {
		return cpu.registers[arg.value]
	} else {
		return arg.value
	}
}

func (cpu *AsmBunnyCPU) incrementPc(cnt int) bool {
	cpu.pc += cnt
	return cpu.pc >= 0 && cpu.pc < len(cpu.prog)
}

func (cpu *AsmBunnyCPU) executeCpy(argA, argB AsmBunnyArgument) bool {
	if argB.isRegister() {
		cpu.registers[argB.value] = cpu.resolveValue(argA)
	}
	return cpu.incrementPc(1)
}

func (cpu *AsmBunnyCPU) executeJnz(argA, argB AsmBunnyArgument) bool {
	if cpu.resolveValue(argA) != 0 {
		return cpu.incrementPc(cpu.resolveValue(argB))
	}
	return cpu.incrementPc(1)
}

func (cpu *AsmBunnyCPU) executeInc(argA AsmBunnyArgument) bool {
	if argA.isRegister() {
		cpu.registers[argA.value]++
	}
	return cpu.incrementPc(1)
}

func (cpu *AsmBunnyCPU) executeDec(argA AsmBunnyArgument) bool {
	if argA.isRegister() {
		cpu.registers[argA.value]--
	}
	return cpu.incrementPc(1)
}

func (cpu *AsmBunnyCPU) executeOut(argA AsmBunnyArgument) bool {
	if cpu.Out != nil {
		cpu.Out <- cpu.resolveValue(argA)
		cpu.outCnt++
		if cpu.outCnt > 10 {
			close(cpu.Out)
			return false
		}
	}
	return cpu.incrementPc(1)
}

func (cpu *AsmBunnyCPU) executeTgl(argA AsmBunnyArgument) bool {
	idx := cpu.pc + cpu.resolveValue(argA)
	if idx >= 0 && idx < len(cpu.prog) {
		i := &cpu.prog[idx]
		switch i.op {
		case INC:
			i.op = DEC
		case DEC:
			i.op = INC
		case TGL:
			i.op = INC
		case OUT:
			i.op = INC
		case CPY:
			i.op = JNZ
		case JNZ:
			i.op = CPY
		}
	}
	return cpu.incrementPc(1)
}

func (cpu *AsmBunnyCPU) RunProgram(prog []AsmBunnyInstruction, a, b, c, d int) []int {
	cpu.registers = make([]int, 0, 4)
	cpu.registers = append(cpu.registers, a, b, c, d)
	cpu.pc = 0
	cpu.prog = prog
	running := true
	for running {
		i := &cpu.prog[cpu.pc]
		switch i.op {
		case CPY:
			running = cpu.executeCpy(i.a, i.b)
		case JNZ:
			running = cpu.executeJnz(i.a, i.b)
		case INC:
			running = cpu.executeInc(i.a)
		case DEC:
			running = cpu.executeDec(i.a)
		case TGL:
			running = cpu.executeTgl(i.a)
		case OUT:
			running = cpu.executeOut(i.a)
		}
	}
	return cpu.registers
}

type AsmBunnyArgument struct {
	value int
	reg   bool
}

func (a *AsmBunnyArgument) asString() string {
	if a.isRegister() {
		return fmt.Sprintf("R(%d)", a.value)
	} else {
		return fmt.Sprintf("L(%d)", a.value)
	}
}

func (arg *AsmBunnyArgument) isRegister() bool {
	return arg.reg
}

func registerArgument(reg byte) AsmBunnyArgument {
	return AsmBunnyArgument{value: int(reg - 'a'), reg: true}
}

func literalArgument(lit int) AsmBunnyArgument {
	return AsmBunnyArgument{value: lit, reg: false}
}

func parseArgument(s string) AsmBunnyArgument {
	if unicode.IsLetter(rune(s[0])) {
		return registerArgument(s[0])
	} else {
		return literalArgument(utils.ReadInt(s))
	}
}

type AsmBunnyInstruction struct {
	op   AsmBunnyOp
	a, b AsmBunnyArgument
}

func (i *AsmBunnyInstruction) asString() string {
	switch i.op {
	case CPY:
		return fmt.Sprintf("CPY\t%s\t%s", i.a.asString(), i.b.asString())
	case JNZ:
		return fmt.Sprintf("JNZ\t%s\t%s", i.a.asString(), i.b.asString())
	case INC:
		return fmt.Sprintf("INC\t%s", i.a.asString())
	case DEC:
		return fmt.Sprintf("DEC\t%s", i.a.asString())
	case TGL:
		return fmt.Sprintf("TGL\t%s", i.a.asString())
	case OUT:
		return fmt.Sprintf("OUT\t%s", i.a.asString())
	}
	return ""
}

type AsmBunnyProgram []AsmBunnyInstruction

func ReadAsmBunnyProgram(fp *bufio.Reader) AsmBunnyProgram {
	var prog []AsmBunnyInstruction
	utils.ReadStrings(fp, func(s string) {
		parts := strings.Split(s, " ")
		switch parts[0] {
		case "cpy":
			prog = append(prog, AsmBunnyInstruction{
				op: CPY,
				a:  parseArgument(parts[1]),
				b:  parseArgument(parts[2])})
		case "jnz":
			prog = append(prog, AsmBunnyInstruction{
				op: JNZ,
				a:  parseArgument(parts[1]),
				b:  parseArgument(parts[2])})
		case "inc":
			prog = append(prog, AsmBunnyInstruction{
				op: INC,
				a:  parseArgument(parts[1])})
		case "dec":
			prog = append(prog, AsmBunnyInstruction{
				op: DEC,
				a:  parseArgument(parts[1])})
		case "tgl":
			prog = append(prog, AsmBunnyInstruction{
				op: TGL,
				a:  parseArgument(parts[1])})
		case "out":
			prog = append(prog, AsmBunnyInstruction{
				op: OUT,
				a:  parseArgument(parts[1])})
		}
	})
	return prog
}

package cpu

import (
	"bufio"
	"strings"
	"unicode"

	"github.com/biesnecker/godvent/types/deque"
	"github.com/biesnecker/godvent/utils"
)

type opcode int

const (
	SND opcode = iota
	SET
	ADD
	MUL
	MOD
	RCV
	JGZ
	JNZ
	SUB
)

type Arg struct {
	value     int
	isLiteral bool
}

func registerArg(reg byte) Arg {
	return Arg{value: int(reg - 'a'), isLiteral: false}
}

func literalArg(value int) Arg {
	return Arg{value: value, isLiteral: true}
}

type Instr struct {
	op   opcode
	a, b Arg
}

func readArgument(arg string) Arg {
	if unicode.IsLetter(rune(arg[0])) {
		return registerArg(arg[0])
	} else {
		return literalArg(utils.ReadInt(arg))
	}
}

type Cpu struct {
	pc               int
	prog             []Instr
	registers        map[int]int
	inputQ, outputQ  *deque.Deque
	sendCount        int
	instructionCount map[opcode]int
}

func NewCpuWithQueues(prog []Instr, inputQ, outputQ *deque.Deque) Cpu {
	r := make(map[int]int)
	return Cpu{
		pc:               0,
		prog:             prog,
		registers:        r,
		inputQ:           inputQ,
		outputQ:          outputQ,
		instructionCount: make(map[opcode]int)}
}

func NewCpu(prog []Instr) Cpu {
	return NewCpuWithQueues(prog, nil, nil)
}

func (c *Cpu) SetRegister(reg byte, value int) {
	c.registers[int(reg-'a')] = value
}

func (c *Cpu) GetRegister(reg byte) int {
	return c.registers[int(reg-'a')]
}

func (c *Cpu) GetSendCount() int {
	return c.sendCount
}

func (c *Cpu) GetInstructionCount(op opcode) int {
	return c.instructionCount[op]
}

func (c *Cpu) resolveValue(arg Arg) int {
	if arg.isLiteral {
		return arg.value
	} else {
		return c.registers[arg.value]
	}
}

func ReadInputProgram(fp *bufio.Reader) []Instr {
	var res []Instr
	utils.ReadStrings(fp, func(s string) {
		parts := strings.Split(s, " ")
		switch parts[0] {
		case "snd":
			res = append(res, Instr{op: SND, a: readArgument(parts[1])})
		case "set":
			res = append(res, Instr{
				op: SET,
				a:  readArgument(parts[1]),
				b:  readArgument(parts[2])})
		case "add":
			res = append(res, Instr{
				op: ADD,
				a:  readArgument(parts[1]),
				b:  readArgument(parts[2])})
		case "mul":
			res = append(res, Instr{
				op: MUL,
				a:  readArgument(parts[1]),
				b:  readArgument(parts[2])})
		case "mod":
			res = append(res, Instr{
				op: MOD,
				a:  readArgument(parts[1]),
				b:  readArgument(parts[2])})
		case "rcv":
			res = append(res, Instr{op: RCV, a: readArgument(parts[1])})
		case "jgz":
			res = append(res, Instr{
				op: JGZ,
				a:  readArgument(parts[1]),
				b:  readArgument(parts[2])})
		case "jnz":
			res = append(res, Instr{
				op: JNZ,
				a:  readArgument(parts[1]),
				b:  readArgument(parts[2])})
		case "sub":
			res = append(res, Instr{
				op: SUB,
				a:  readArgument(parts[1]),
				b:  readArgument(parts[2])})
		}
	})

	return res
}

func (c *Cpu) ExecuteProgram(partA bool) (int, bool) {
	madeProgress := false
	for c.pc >= 0 && c.pc < len(c.prog) {
		i := &c.prog[c.pc]
		switch i.op {
		case SND:
			c.outputQ.PushRight(c.resolveValue(i.a))
			c.sendCount++
			c.instructionCount[SND]++
			madeProgress = true
			c.pc++
		case SET:
			c.registers[i.a.value] = c.resolveValue(i.b)
			c.instructionCount[SET]++
			madeProgress = true
			c.pc++
		case ADD:
			c.registers[i.a.value] += c.resolveValue(i.b)
			c.instructionCount[ADD]++
			madeProgress = true
			c.pc++
		case SUB:
			c.registers[i.a.value] -= c.resolveValue(i.b)
			c.instructionCount[SUB]++
			madeProgress = true
			c.pc++
		case MUL:
			c.registers[i.a.value] *= c.resolveValue(i.b)
			c.instructionCount[MUL]++
			madeProgress = true
			c.pc++
		case MOD:
			c.registers[i.a.value] %= c.resolveValue(i.b)
			c.instructionCount[MOD]++
			madeProgress = true
			c.pc++
		case RCV:
			if partA {
				if c.resolveValue(i.a) != 0 {
					return c.inputQ.PopRight().(int), false
				}
			} else {
				if c.inputQ.Empty() {
					return 0, madeProgress
				} else {
					c.registers[i.a.value] = c.inputQ.PopLeft().(int)
				}
			}
			c.instructionCount[RCV]++
			madeProgress = true
			c.pc++
		case JGZ:
			if c.resolveValue(i.a) > 0 {
				c.pc += c.resolveValue(i.b)
			} else {
				c.pc++
			}
			madeProgress = true
			c.instructionCount[JGZ]++
		case JNZ:
			if c.resolveValue(i.a) != 0 {
				c.pc += c.resolveValue(i.b)
			} else {
				c.pc++
			}
			madeProgress = true
			c.instructionCount[JNZ]++
		}
	}
	return 0, madeProgress
}

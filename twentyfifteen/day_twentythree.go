package twentyfifteen

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

const (
	HLF int = 0
	TPL int = 1
	INC int = 2
	JMP int = 3
	JIE int = 4
	JIO int = 5
)

type instr struct {
	op     int
	reg    byte
	offset int
}

func readInputDayTwentyThree(fp *bufio.Reader) []instr {
	var prog []instr
	utils.ReadStrings(fp, func(s string) {
		s = strings.ReplaceAll(s, ", ", " ")
		parts := strings.Split(s, " ")
		switch parts[0] {
		case "hlf":
			prog = append(prog, instr{op: HLF, reg: parts[1][0]})
		case "tpl":
			prog = append(prog, instr{op: TPL, reg: parts[1][0]})
		case "inc":
			prog = append(prog, instr{op: INC, reg: parts[1][0]})
		case "jmp":
			if offset, err := strconv.Atoi(parts[1]); err != nil {
				log.Fatalln(err)
			} else {
				prog = append(prog, instr{op: JMP, offset: offset})
			}
		case "jie":
			if offset, err := strconv.Atoi(parts[2]); err != nil {
				log.Fatalln(err)
			} else {
				prog = append(prog, instr{op: JIE, reg: parts[1][0], offset: offset})
			}
		case "jio":
			if offset, err := strconv.Atoi(parts[2]); err != nil {
				log.Fatalln(err)
			} else {
				prog = append(prog, instr{op: JIO, reg: parts[1][0], offset: offset})
			}
		}
	})
	return prog
}

func runProgram(a, b int, is []instr) (int, int) {
	var pc int
	for pc < len(is) {
		switch is[pc].op {
		case HLF:
			if is[pc].reg == 'a' {
				a /= 2
			} else {
				b /= 2
			}
			pc++
		case TPL:
			if is[pc].reg == 'a' {
				a *= 3
			} else {
				b *= 3
			}
			pc++
		case INC:
			if is[pc].reg == 'a' {
				a++
			} else {
				b++
			}
			pc++
		case JMP:
			pc += is[pc].offset
		case JIE:
			if (is[pc].reg == 'a' && a%2 == 0) || (is[pc].reg == 'b' && b%2 == 0) {
				pc += is[pc].offset
			} else {
				pc++
			}
		case JIO:
			if (is[pc].reg == 'a' && a == 1) || (is[pc].reg == 'b' && b == 1) {
				pc += is[pc].offset
			} else {
				pc++
			}
		}
	}
	return a, b
}

func DayTwentyThreeA(fp *bufio.Reader) string {
	_, b := runProgram(0, 0, readInputDayTwentyThree(fp))
	return strconv.Itoa(b)
}

func DayTwentyThreeB(fp *bufio.Reader) string {
	_, b := runProgram(1, 0, readInputDayTwentyThree(fp))
	return strconv.Itoa(b)
}

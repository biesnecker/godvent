package twentyfifteen

import (
	"bufio"
	"log"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/utils"
)

type logicgate struct {
	source   []string
	value    int
	resolved bool
}

func parseInputDaySeven(gates *map[string]*logicgate) func(string) {
	return func(s string) {
		parts := strings.Fields(s)
		tail := parts[len(parts)-1]

		(*gates)[tail] = &logicgate{source: parts, value: 0, resolved: false}
	}
}

func resolveIdentifierOrLiteral(gates map[string]*logicgate, val string) int {
	if intval, err := strconv.Atoi(val); err != nil {
		return resolveGate(gates, val)
	} else {
		return intval
	}
}

func resolveGate(gates map[string]*logicgate, name string) int {
	gate := gates[name]
	returnValue := 0
	if gate.resolved {
		return gate.value
	}

	if gate.source[0] == "NOT" {
		val := resolveGate(gates, gate.source[1])
		returnValue = ^val
	} else if gate.source[1] == "->" {
		returnValue = resolveIdentifierOrLiteral(gates, gate.source[0])
	} else if gate.source[1] == "OR" {
		valOne := resolveIdentifierOrLiteral(gates, gate.source[0])
		valTwo := resolveIdentifierOrLiteral(gates, gate.source[2])
		returnValue = valOne | valTwo
	} else if gate.source[1] == "AND" {
		valOne := resolveIdentifierOrLiteral(gates, gate.source[0])
		valTwo := resolveIdentifierOrLiteral(gates, gate.source[2])
		returnValue = valOne & valTwo
	} else if gate.source[1] == "LSHIFT" {
		val := resolveIdentifierOrLiteral(gates, gate.source[0])
		if amount, err := strconv.Atoi(gate.source[2]); err != nil {
			log.Fatal(err)
		} else {
			returnValue = val << amount
		}
	} else if gate.source[1] == "RSHIFT" {
		val := resolveIdentifierOrLiteral(gates, gate.source[0])
		if amount, err := strconv.Atoi(gate.source[2]); err != nil {
			log.Fatal(err)
		} else {
			returnValue = val >> amount
		}
	} else {
		log.Fatalf("Unknown command: %v\n", gate.source)
	}
	gate.value = returnValue
	gate.resolved = true
	return returnValue
}

func DaySevenA(fp *bufio.Reader) string {
	gates := make(map[string]*logicgate)
	utils.ReadStrings(fp, parseInputDaySeven(&gates))
	solution := resolveGate(gates, "a")
	return strconv.Itoa(solution)
}

func DaySevenB(fp *bufio.Reader) string {
	gates := make(map[string]*logicgate)
	utils.ReadStrings(fp, parseInputDaySeven(&gates))

	// Hack the value of b to be the previous answer, 3176
	bgate := gates["b"]
	bgate.resolved = true
	bgate.value = 3176

	solution := resolveGate(gates, "a")
	return strconv.Itoa(solution)
}

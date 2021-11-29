package twentysixteen

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/biesnecker/godvent/twentysixteen/dayeleven"
	"github.com/biesnecker/godvent/utils"
)

func readInputDayEleven(fp *bufio.Reader) dayeleven.RTGState {
	state := dayeleven.RTGState{}

	elementTypes := make(map[string]int)
	elementIdx := int(0)

	getElementId := func(name string) int {
		if eid, found := elementTypes[name]; !found {
			elementTypes[name] = elementIdx
			eid = elementIdx
			elementIdx++
			return eid
		} else {
			return eid
		}
	}

	utils.ReadStringsWithIndex(fp, func(floor int, s string) {
		if strings.Contains(s, "nothing") {
			return
		}
		parts := strings.Split(s, " ")
		for i, part := range parts {
			if strings.HasPrefix(part, "generator") {
				state.SetGenerator(floor, getElementId(parts[i-1]))
			} else if strings.HasSuffix(part, "compatible") {
				state.SetMicrochip(floor,
					getElementId(
						parts[i][:strings.IndexByte(parts[i], '-')]))
			}
		}
	})
	state.SetElevator(0)
	return state
}

func simulateDayEleven(initialState dayeleven.RTGState) int {
	q := make(chan *dayeleven.RTGPosition, 500000)

	q <- dayeleven.NewRTGPosition(initialState, 0)

	seen := make(map[dayeleven.RTGStateRepr]struct{})

complete:
	for {
		select {
		case p := <-q:
			if p.IsComplete() {
				return p.GetSteps()
			}
			p.GenerateNextPositions(func(newPosition *dayeleven.RTGPosition) {
				s := newPosition.GetState().GetRepr()
				if _, found := seen[s]; !found {
					seen[s] = struct{}{}
					q <- newPosition
				}
			})
		default:
			break complete
		}
	}
	return 0
}

func DayElevenA(fp *bufio.Reader) string {
	s := readInputDayEleven(fp)
	return strconv.Itoa(simulateDayEleven(s))
}

func DayElevenB(fp *bufio.Reader) string {
	s := readInputDayEleven(fp)
	s.SetGenerator(0, 5)
	s.SetGenerator(0, 6)
	s.SetMicrochip(0, 5)
	s.SetMicrochip(0, 6)
	return strconv.Itoa(simulateDayEleven(s))
}

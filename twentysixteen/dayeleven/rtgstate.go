package dayeleven

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type RTGState struct {
	generators [4]uint8
	microchips [4]uint8
	elevator   int
}

type RTGStateRepr struct {
	data uint64
}

func (r *RTGState) GetRepr() RTGStateRepr {
	type reprelem struct {
		mpos, gpos uint8
	}

	elems := make([]reprelem, 8)

	for i := 0; i < 4; i++ {
		m := r.microchips[i]
		g := r.generators[i]
		for j := 0; j < 8; j++ {
			check := uint8(1) << j
			if m&check > 0 {
				elems[j].mpos = uint8(i)
			}
			if g&check > 0 {
				elems[j].gpos = uint8(i)
			}
		}
	}
	sort.Slice(elems, func(i, j int) bool {
		if elems[i].mpos == elems[j].mpos {
			return elems[i].gpos < elems[j].gpos
		} else {
			return elems[i].mpos < elems[j].mpos
		}
	})

	var state uint64
	for i := range elems {
		state <<= 2
		state |= uint64(elems[i].mpos & 0x3)
		state <<= 2
		state |= uint64(elems[i].gpos & 0x3)
	}

	state <<= 2
	state |= uint64(r.elevator & 0x3)

	return RTGStateRepr{state}
}

func mask(pos int) uint8 {
	return 1 << pos
}

func (r *RTGState) GetElevator() int {
	return r.elevator
}

func (r *RTGState) SetElevator(pos int) {
	r.elevator = pos
}

func (r *RTGState) IsValid() bool {
	for i := 0; i < 4; i++ {
		if r.generators[i] != 0 && r.microchips[i]&(^r.generators[i]) != 0 {
			return false
		}
	}
	return true
}

func (r *RTGState) HasElement(floor, pos int) bool {
	if pos < 8 {
		return r.HasGenerator(floor, pos)
	} else {
		return r.HasMicrochip(floor, pos-8)
	}
}

func (r *RTGState) SetElement(floor, pos int) {
	if pos < 8 {
		r.SetGenerator(floor, pos)
	} else {
		r.SetMicrochip(floor, pos-8)
	}
}

func (r *RTGState) UnsetElement(floor, pos int) {
	if pos < 8 {
		r.UnsetGenerator(floor, pos)
	} else {
		r.UnsetMicrochip(floor, pos-8)
	}
}

func (r *RTGState) MoveElement(startFloor, endFloor, pos int) {
	if pos < 8 {
		r.MoveGenerator(startFloor, endFloor, pos)
	} else {
		r.MoveMicrochip(startFloor, endFloor, pos-8)
	}
}

func (r *RTGState) HasMicrochip(floor, pos int) bool {
	return r.microchips[floor]&mask(pos) > 0
}

func (r *RTGState) HasGenerator(floor, pos int) bool {
	return r.generators[floor]&mask(pos) > 0
}

func (r *RTGState) SetMicrochip(floor, pos int) {
	r.microchips[floor] |= mask(pos)
}

func (r *RTGState) SetGenerator(floor, pos int) {
	r.generators[floor] |= mask(pos)
}

func (r *RTGState) UnsetMicrochip(floor, pos int) {
	r.microchips[floor] &= ^mask(pos)
}

func (r *RTGState) UnsetGenerator(floor, pos int) {
	r.generators[floor] &= ^mask(pos)
}

func (r *RTGState) MoveMicrochip(startFloor, endFloor, pos int) {
	r.UnsetMicrochip(startFloor, pos)
	r.SetMicrochip(endFloor, pos)
}

func (r *RTGState) MoveGenerator(startFloor, endFloor, pos int) {
	r.UnsetGenerator(startFloor, pos)
	r.SetGenerator(endFloor, pos)
}

func (r *RTGState) IsFloorEmpty(floor int) bool {
	return r.microchips[floor] == 0 && r.generators[floor] == 0
}

func (r *RTGState) IsComplete() bool {
	for i := 0; i < 3; i++ {
		if r.microchips[i] != 0 || r.generators[i] != 0 {
			return false
		}
	}
	return true
}

func (r *RTGState) PrintBinaryDebug() string {
	lines := make([]string, 0, 5)
	lines = append(lines, "     Chips     Gens")

	for i := 3; i >= 0; i-- {
		parts := make([]string, 0, 4)
		parts = append(parts, fmt.Sprintf("%d  ", i))
		if r.GetElevator() == i {
			parts = append(parts, "E ")
		} else {
			parts = append(parts, "  ")
		}
		parts = append(parts, fmt.Sprintf("%08s  ",
			strconv.FormatUint(uint64(r.microchips[i]), 2)))
		parts = append(parts, fmt.Sprintf("%08s",
			strconv.FormatUint(uint64(r.generators[i]), 2)))
		lines = append(lines, strings.Join(parts, ""))
	}

	return strings.Join(lines, "\n")
}

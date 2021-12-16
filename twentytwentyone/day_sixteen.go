package twentytwentyone

import (
	"bufio"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

func numberFromBits(bits []bool) uint64 {
	n := uint64(0)
	for _, b := range bits {
		n <<= 1
		if b {
			n |= 1
		}
	}
	return n
}

func decodeLiteral(bits []bool) (int, uint64) {
	res := uint64(0)
	consumed := 0
	stop := false
	for !stop {
		stop = !bits[0]
		bits = bits[1:]
		for i := 0; i < 4; i++ {
			res <<= 1
			if bits[i] {
				res |= 1
			}
		}
		bits = bits[4:]
		consumed += 5
	}
	return consumed, res
}

func readInputD16(fp *bufio.Reader) []bool {
	res := make([]bool, 0, 128)
	s := utils.ReadSingleString(fp)
	for j := range s {
		switch s[j] {
		case '0':
			res = append(res, false, false, false, false)
		case '1':
			res = append(res, false, false, false, true)
		case '2':
			res = append(res, false, false, true, false)
		case '3':
			res = append(res, false, false, true, true)
		case '4':
			res = append(res, false, true, false, false)
		case '5':
			res = append(res, false, true, false, true)
		case '6':
			res = append(res, false, true, true, false)
		case '7':
			res = append(res, false, true, true, true)
		case '8':
			res = append(res, true, false, false, false)
		case '9':
			res = append(res, true, false, false, true)
		case 'A':
			res = append(res, true, false, true, false)
		case 'B':
			res = append(res, true, false, true, true)
		case 'C':
			res = append(res, true, true, false, false)
		case 'D':
			res = append(res, true, true, false, true)
		case 'E':
			res = append(res, true, true, true, false)
		case 'F':
			res = append(res, true, true, true, true)
		}
	}
	return res
}

type packet struct {
	version  uint64
	typeId   uint64
	value    uint64
	children []*packet
}

func parseBinaryString(bits []bool, pad bool) (int, *packet) {
	if len(bits) == 0 {
		return 0, nil
	}
	version := numberFromBits(bits[:3])
	typeId := numberFromBits(bits[3:6])
	bits = bits[6:]
	p := &packet{version: version, typeId: typeId}
	consumed := 6
	switch typeId {
	case 4:
		var lit uint64
		var nc int
		nc, lit = decodeLiteral(bits)
		consumed += nc
		p.value = lit
	default:
		optype := bits[0]
		bits = bits[1:]
		consumed++
		switch optype {
		case true:
			consumed += 11
			nsubs := int(numberFromBits(bits[:11]))
			bits = bits[11:]
			for i := 0; i < nsubs; i++ {
				nc, child := parseBinaryString(bits, false)
				p.children = append(p.children, child)
				consumed += nc
				bits = bits[nc:]
			}
		case false:
			consumed += 15
			sublen := int(numberFromBits(bits[:15]))
			bits = bits[15:]
			for sublen > 0 {
				nc, child := parseBinaryString(bits, false)
				p.children = append(p.children, child)
				consumed += nc
				sublen -= nc
				bits = bits[nc:]
			}
		}
		switch typeId {
		case 0:
			for _, c := range p.children {
				p.value += c.value
			}
		case 1:
			product := uint64(1)
			for _, c := range p.children {
				product *= c.value
			}
			p.value = product
		case 2:
			min := uint64(math.MaxUint64)
			for _, c := range p.children {
				if c.value < min {
					min = c.value
				}
			}
			p.value = min
		case 3:
			max := uint64(0)
			for _, c := range p.children {
				if c.value > max {
					max = c.value
				}
			}
			p.value = max
		case 5:
			if p.children[0].value > p.children[1].value {
				p.value = 1
			} else {
				p.value = 0
			}
		case 6:
			if p.children[0].value < p.children[1].value {
				p.value = 1
			} else {
				p.value = 0
			}
		case 7:
			if p.children[0].value == p.children[1].value {
				p.value = 1
			} else {
				p.value = 0
			}
		}
	}
	if pad && consumed%4 != 0 {
		consumed += (4 - (consumed % 4))
	}
	return consumed, p
}

func sumPacketVersion(p *packet) uint64 {
	total := p.version
	for _, c := range p.children {
		total += sumPacketVersion(c)
	}
	return total
}

func DaySixteenA(fp *bufio.Reader) string {
	input := readInputD16(fp)
	_, p := parseBinaryString(input, false)
	return strconv.Itoa(int(sumPacketVersion(p)))
}

func DaySixteenB(fp *bufio.Reader) string {
	input := readInputD16(fp)
	_, p := parseBinaryString(input, false)
	return strconv.Itoa(int(p.value))
}

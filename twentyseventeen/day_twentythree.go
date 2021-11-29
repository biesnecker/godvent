package twentyseventeen

import (
	"bufio"
	"strconv"

	"github.com/biesnecker/godvent/twentyseventeen/cpu"
)

func DayTwentyThreeA(fp *bufio.Reader) string {
	prog := cpu.ReadInputProgram(fp)
	c := cpu.NewCpu(prog)

	c.ExecuteProgram(false)
	return strconv.Itoa(c.GetInstructionCount(cpu.MUL))
}

func DayTwentyThreeB(fp *bufio.Reader) string {
	/*
		let r = {
		    b: 67,
		    c: 67,
		    d: 0,
		    f: 0,
		    g: 0,
		    h: 0
		  }
		  r['b'] = r['b'] * 100 + 100000
		  r['c'] = r['b'] + 17000
		  do {
		    r['f'] = 1
		    r['d'] = 2
		    for (let d = r['d']; d * d < r['b']; ++d) {
		      if (r['b'] % d === 0) {
		        r['f'] = 0
		        break
		      }
		    }
		    if (r['f'] === 0) r['h']++
		    r['g'] = r['b'] - r['c']
		    r['b'] += 17
		  } while (r['g'] !== 0)

		  return r['h']

	*/

	r := map[byte]int{
		'a': 1, 'b': 81, 'c': 81, 'd': 0,
		'e': 0, 'f': 0, 'g': 0, 'h': 0}

	r['b'] = r['b']*100 + 100000
	r['c'] = r['b'] + 17000
	for {
		r['f'] = 1
		r['d'] = 2
		for d := r['d']; d*d < r['b']; d++ {
			if r['b']%d == 0 {
				r['f'] = 0
				break
			}
		}
		if r['f'] == 0 {
			r['h']++
		}
		r['g'] = r['b'] - r['c']
		r['b'] += 17
		if r['g'] == 0 {
			break
		}
	}

	return strconv.Itoa(r['h'])
}

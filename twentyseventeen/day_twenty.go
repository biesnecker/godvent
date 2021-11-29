package twentyseventeen

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/biesnecker/godvent/types"
	"github.com/biesnecker/godvent/utils"
)

type particle struct {
	id              int
	p, v, a         types.Coord3
	distance, delta int
	collided        bool
}

func (p *particle) tick() {
	// Update velocity
	p.v.X += p.a.X
	p.v.Y += p.a.Y
	p.v.Z += p.a.Z
	// Update position
	p.p.X += p.v.X
	p.p.Y += p.v.Y
	p.p.Z += p.v.Z
	oldDistance := p.distance
	//oldDelta := p.delta
	p.distance = utils.ManhattanDistance3(p.p, types.Coord3{X: 0, Y: 0, Z: 0})
	p.delta = p.distance - oldDistance
}

func readInputDayTwenty(fp *bufio.Reader) []particle {
	var res []particle
	utils.ReadStringsWithIndex(fp, func(i int, s string) {
		var px, py, pz, vx, vy, vz, ax, ay, az int
		fmt.Sscanf(s,
			"p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>",
			&px, &py, &pz, &vx, &vy, &vz, &ax, &ay, &az)
		particle := particle{
			id: i,
			p:  types.Coord3{X: px, Y: py, Z: pz},
			v:  types.Coord3{X: vx, Y: vy, Z: vz},
			a:  types.Coord3{X: ax, Y: ay, Z: az}}
		particle.distance = utils.ManhattanDistance3(
			particle.p,
			types.Coord3{X: 0, Y: 0, Z: 0})
		res = append(res, particle)
	})
	return res
}

func DayTwentyA(fp *bufio.Reader) string {
	particles := readInputDayTwenty(fp)
	for i := 0; i < 1000; i++ {
		for j := range particles {
			particles[j].tick()
		}
	}
	minDistance := math.MaxInt64
	minDistanceId := 0
	for i := range particles {
		if particles[i].distance < minDistance {
			minDistance = particles[i].distance
			minDistanceId = particles[i].id
		}
	}
	return strconv.Itoa(minDistanceId)
}

func DayTwentyB(fp *bufio.Reader) string {
	particles := readInputDayTwenty(fp)
	for i := 0; i < 1000; i++ {
		collisions := make(map[types.Coord3]int)
		for j := range particles {
			if !particles[j].collided {
				particles[j].tick()
				if otherP, found := collisions[particles[j].p]; found {
					particles[otherP].collided = true
					particles[j].collided = true
				} else {
					collisions[particles[j].p] = particles[j].id
				}
			}
		}
	}
	count := 0
	for i := range particles {
		if !particles[i].collided {
			count++
		}
	}
	return strconv.Itoa(count)
}

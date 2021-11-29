package twentyfifteen

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/biesnecker/godvent/utils"
)

type reindeerStats struct {
	speed       int
	flyingTime  int
	restingTime int
	period      int
}

func readInputDayFourteen(fp *bufio.Reader) []*reindeerStats {
	stats := make([]*reindeerStats, 0, 10)
	utils.ReadStrings(fp, func(s string) {
		var reindeerName string
		var speed, flyingTime, restingTime int
		fmt.Sscanf(
			s,
			"%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&reindeerName, &speed, &flyingTime, &restingTime)

		stats = append(
			stats,
			&reindeerStats{
				speed:       speed,
				flyingTime:  flyingTime,
				restingTime: restingTime,
				period:      flyingTime + restingTime})
	})
	return stats
}

func DayFourteenA(fp *bufio.Reader) string {
	stats := readInputDayFourteen(fp)
	distances := make([]int, len(stats))

	for s := 0; s < 2503; s++ {
		for rid, rs := range stats {
			if (s % rs.period) < rs.flyingTime {
				distances[rid] += rs.speed
			}
		}
	}

	maxDistance := 0
	for _, d := range distances {
		if d > maxDistance {
			maxDistance = d
		}
	}
	return strconv.Itoa(maxDistance)
}

func DayFourteenB(fp *bufio.Reader) string {
	stats := readInputDayFourteen(fp)
	distances := make([]int, len(stats))
	points := make([]int, len(stats))

	for s := 0; s < 2503; s++ {
		for rid, rs := range stats {
			if (s % rs.period) < rs.flyingTime {
				distances[rid] += rs.speed
			}
		}
		leader := -1
		maxDistance := 0
		for i, d := range distances {
			if d > maxDistance {
				maxDistance = d
				leader = i
			}
		}
		points[leader]++
	}

	maxPoints := 0
	for _, p := range points {
		if p > maxPoints {
			maxPoints = p
		}
	}
	return strconv.Itoa(maxPoints)
}

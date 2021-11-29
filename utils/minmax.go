package utils

import (
	"log"
	"math"
)

func BoundInt(x, min, max int) int {
	if x < min {
		return min
	} else if x > max {
		return max
	} else {
		return x
	}
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinIntSlice(s []int) int {
	if len(s) == 0 {
		log.Fatalln("Zero-length slice passed to MinIntSlice.")
	}
	m := math.MaxInt64
	for _, v := range s {
		if v < m {
			m = v
		}
	}
	return m
}

func MaxIntSlice(s []int) int {
	if len(s) == 0 {
		log.Fatalln("Zero-length slice passed to MaxIntSlice.")
	}
	m := math.MinInt64
	for _, v := range s {
		if v > m {
			m = v
		}
	}
	return m
}

func MinMaxIntSlice(s []int) (int, int) {
	if len(s) == 0 {
		log.Fatalln("Zero-length slice passed to MinMaxIntSlice")
	}
	min := s[0]
	max := s[0]
	for _, v := range s[1:] {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min, max
}

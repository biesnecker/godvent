package utils

func SumOfIntSlice(ns []int) int {
	sum := 0
	for i := range ns {
		sum += ns[i]
	}
	return sum
}

func ProductOfIntSlice(ns []int) int {
	prod := 1
	for i := range ns {
		prod *= ns[i]
	}
	return prod
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IntPow(base, exp int) int {
	b := base
	for i := 1; i < exp; i++ {
		b *= base
	}
	return b
}

package utils

func Combinations(n, r int, handler func([]int) bool) {
	if r >= n {
		panic("r cannot exceed n")
	}
	combo := make([]int, r)
	for i := 0; i < r; i++ {
		combo[i] = i
	}
	i := r - 1
	for combo[0] < n-r+1 {
		for i > 0 && combo[i] == n-r+i {
			i--
		}

		if !handler(combo) {
			break
		}

		combo[i]++
		for i < r-1 {
			combo[i+1] = combo[i] + 1
			i++
		}
	}
}

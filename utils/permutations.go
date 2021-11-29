package utils

func Permutations(n int, handler func([]int) bool) {
	if n <= 0 {
		panic("n can't be less than zero")
	}
	if n == 1 {
		handler([]int{0})
		return
	}

	perm := make([]int, n)
	indexes := make([]int, n)

	for i := 0; i < n; i++ {
		perm[i] = i
	}

	if !handler(perm) {
		return
	}

	i := 1
	for i < n {
		if indexes[i] < i {
			if i&1 == 1 {
				perm[i], perm[indexes[i]] = perm[indexes[i]], perm[i]
			} else {
				perm[i], perm[0] = perm[0], perm[i]
			}
			if !handler(perm) {
				return
			}
			indexes[i]++
			i = 1
		} else {
			indexes[i] = 0
			i++
		}
	}
}

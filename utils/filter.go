package utils

func FilterInts(nums []int, filterF func(int) bool) []int {
	var res []int
	for _, v := range nums {
		if filterF(v) {
			res = append(res, v)
		}
	}
	return res
}

func FilterStrings(strings []string, filterF func(string) bool) []string {
	var res []string
	for _, s := range strings {
		if filterF(s) {
			res = append(res, s)
		}
	}
	return res
}

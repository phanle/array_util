package array_util

func Min(values ...int) int {
	min := values[0]
	for _, val := range values {
		if val <= min {
			min = val
		}
	}
	return min
}

func MinIn(values []int) int {
	min := values[0]
	for _, val := range values {
		if val <= min {
			min = val
		}
	}
	return min
}

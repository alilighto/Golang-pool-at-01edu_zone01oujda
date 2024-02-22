package piscine

func Max(a []int) int {
	max := 0
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

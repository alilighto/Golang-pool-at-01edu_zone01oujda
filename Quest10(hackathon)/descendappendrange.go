package piscine

func DescendAppendRange(max, min int) []int {
	var s []int
	if max > min {
		for i := max; i > min; i-- {
			s = append(s, i)
		}
	} else {
		return []int{}
	}
	return s
}

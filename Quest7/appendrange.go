package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var table []int
	for i := min; i < max; i++ {
		table = append(table, i)
	}
	return table
}

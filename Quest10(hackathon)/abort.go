package piscine

func Abort(a, b, c, d, e int) int {
	slice := []int{a, b, c, d, e}
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[j] > slice[i] {
				slice[j], slice[i] = slice[i], slice[j]
			}
		}
	}
	return slice[2]
}

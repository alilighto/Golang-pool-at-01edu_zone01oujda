package piscine

func Unmatch(a []int) int {
	for i, nb1 := range a {
		counter := 1
		for j, nb2 := range a {
			if nb1 == nb2 && i != j {
				counter++
			}
		}
		if counter%2 == 1 {
			return nb1
		}
	}
	return -1
}

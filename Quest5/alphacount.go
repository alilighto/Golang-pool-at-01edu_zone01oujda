package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, i := range s {
		if i >= 'A' && i <= 'Z' || i >= 'a' && i <= 'z' {
			counter++
		}
	}
	return counter
}

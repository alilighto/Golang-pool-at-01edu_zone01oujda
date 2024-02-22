package piscine

func NRune(s string, n int) rune {
	va1 := []rune(s)
	if n <= 0 {
		return 0
	}
	if len(va1) == 0 || len(va1) < n {
		return 0
	}
	return va1[n-1]
}

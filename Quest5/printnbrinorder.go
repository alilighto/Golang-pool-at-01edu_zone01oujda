package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var s []int
	for i := 0; n > 0; i++ {
		s = append(s, n%10)
		n /= 10
	}

	l := len(s)
	for i := 1; i < l; i++ {
		if s[i-1] > s[i] {
			s[i], s[i-1] = s[i-1], s[i]
			i = 0 // Reset the index to 0 after a swap to recheck the previous elements
		}
	}

	for i := range s {
		z01.PrintRune(rune(s[i] + '0'))
	}
}

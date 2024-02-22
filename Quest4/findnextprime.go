package piscine

func FindNextPrime(x int) int {
	if x > 1 {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return FindNextPrime(x + 1)
			}
		}
		return x
	} else {
		return 2
	}
}

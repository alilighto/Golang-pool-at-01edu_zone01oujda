package piscine

func RecursivePower(nb, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	return RecursivePower(nb, power-1) * nb
}

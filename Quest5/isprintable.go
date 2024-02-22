package piscine

func IsPrintable(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, i := range str {
		if !(32 <= i && i <= 126) {
			return false
		}
	}
	return true
}

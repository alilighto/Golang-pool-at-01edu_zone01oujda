package piscine

func IsAlpha(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, i := range str {
		if !(('0' <= i && i <= '9') ||
			('A' <= i && i <= 'Z') ||
			('a' <= i && i <= 'z')) {
			return false
		}
	}
	return true
}

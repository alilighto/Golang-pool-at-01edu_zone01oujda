package piscine

func IsNumeric(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, i := range str {
		if !('0' <= i && i <= '9') {
			return false
		}
	}
	return true
}

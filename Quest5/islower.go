package piscine

func IsLower(str string) bool {
	if len(str) > 0 {
		for _, i := range str {
			if !(i >= 'a' && i <= 'z') {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

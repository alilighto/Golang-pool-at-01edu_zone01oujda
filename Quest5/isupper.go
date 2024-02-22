package piscine

func IsUpper(str string) bool {
	if len(str) > 0 {
		for _, i := range str {
			if !(i > 'A' && i < 'Z') {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

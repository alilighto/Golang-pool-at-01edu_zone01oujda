package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for i, v := range s {
		if v >= 'a' && v <= 'z' {
			str[i] = v - 32
		} else {
			str[i] = v
		}
	}
	return string(str)
}

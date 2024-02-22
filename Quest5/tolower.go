package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i, v := range s {
		if v >= 'A' && v <= 'Z' {
			str[i] = v + 32
		} else {
			str[i] = v
		}
	}
	return string(str)
}

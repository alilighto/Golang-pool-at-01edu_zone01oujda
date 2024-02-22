package piscine

func JumpOver(s string) string {
	result := ""
	for i, v := range s {
		if (i+1)%3 == 0 {
			result += string(v)
		}
	}
	return result + "\n"
}

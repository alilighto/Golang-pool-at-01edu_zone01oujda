package piscine

func TrimAtoi(s string) int {
	sign := 1
	var result int
	var i int
	p := true
	for i < len(s) {
		if s[i] == '-' && p == true {
			sign = -1
		} else if s[i] >= 48 && s[i] <= 57 {
			result = result*10 + int(s[i]-48)
			p = false
		}
		i++
	}
	return result * sign
}

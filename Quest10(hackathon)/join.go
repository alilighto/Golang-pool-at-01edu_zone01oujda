package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := strs[0]
	for _, str := range strs[1:] {
		result += sep + str
	}

	return result
}

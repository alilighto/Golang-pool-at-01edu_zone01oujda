package piscine

func SplitWhiteSpaces(s string) []string {
	resault := []string(nil)
	word := ""

	for _, char := range s {
		if char == ' ' {
			if word != "" {
				resault = append(resault, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}

	if word != "" {
		resault = append(resault, word)
	}

	return resault
}

package piscine

func Compact(ptr *[]string) int {
	src := *ptr
	var result []string

	for _, value := range src {
		if value != "" {
			result = append(result, value)
		}
	}

	*ptr = result
	return len(result)
}

package piscine

func ConcatParams(args []string) string {
	var res string
	if len(args) == 0 {
		return ""
	}
	for i, v := range args {
		res += v
		if i < len(args)-1 {
			res += "\n"
		}
	}
	return res
}

package piscine

func StringToIntSlice(str string) []int {
	var arr []int
	for _, v := range str {
		arr = append(arr, int(v))
	}
	return arr
}

package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a / *b
	*b = *a % *b
	*a = c
}

// func main() {
// 	a := 13
// 	b := 2
// 	UltimateDivMod(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

package piscine

// import "fmt"

func ReverseMenuIndex(menu []string) []string {
	a := menu
	for j := len(menu) - 1; j >= 0; j-- {
		a[len(menu)-1-j] = a[j]
	}
	return menu
}

// func main() {
// 	fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
// }

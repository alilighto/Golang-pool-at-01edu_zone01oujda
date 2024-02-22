package piscine

// import "fmt"

func PodiumPosition(podium [][]string) [][]string {
	for a := 0; a < len(podium)-1; a++ {
		for b := a + 1; b < len(podium); b++ {
			if podium[a][0] > podium[b][0] {
				// Swap elements if length is greater
				podium[a], podium[b] = podium[b], podium[a]
			}
		}
	}
	return podium
}

// func main() {
// 	p4 := []string{"4th Place"}
// 	p3 := []string{"3rd Place"}
// 	p2 := []string{"2nd Place"}
// 	p1 := []string{"1st Place"}

// 	position := [][]string{p4, p3, p2, p1}
// 	fmt.Println(PodiumPosition(position))
// }

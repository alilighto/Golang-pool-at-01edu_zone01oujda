package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(table []string) {
	for _, W := range table {
		for _, L := range W {
			z01.PrintRune(L)
		}
		z01.PrintRune('\n')
	}
}

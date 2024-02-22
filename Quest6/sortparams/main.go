package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	length := len(args)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if args[j] > args[j+1] {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	for _, arg := range args {
		for _, i := range arg {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}

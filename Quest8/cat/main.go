package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func put(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func fatal(s string) {
	put(s + "\n")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fatal(err.Error())
		}
		put(string(b))
	} else {
		for _, file := range args {
			b, err := ioutil.ReadFile(file)
			if err != nil {
				fatal("ERROR: " + err.Error())
			}
			put(string(b))
		}
	}
}

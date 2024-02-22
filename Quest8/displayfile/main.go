package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	/* Check if the filename is provided as an argument
	as we know the fist argument is the programme name */
	if len(os.Args) != 2 {
		if len(os.Args) < 2 {
			fmt.Println("File name missing")
		} else {
			fmt.Println("Too many arguments")
		}
		return
	}
	// Get the filename from the command-line arguments
	// Read the content of the file
	content, _ := ioutil.ReadFile(os.Args[1])
	fmt.Print(string(content))
}

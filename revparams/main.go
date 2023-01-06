package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := len(os.Args)
	for i := params - 1; i >= 1; i-- {
		for _, a := range os.Args[i] {
			z01.PrintRune(a)
		}
		z01.PrintRune(10)
	}
}

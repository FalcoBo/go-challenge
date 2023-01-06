package main

import "github.com/01-edu/z01"

func main() {
	for compt := 48; compt < 58; compt++ {
		z01.PrintRune(rune(compt))
	}
	z01.PrintRune(rune('\n'))
}

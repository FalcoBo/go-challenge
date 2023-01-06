package main

import "github.com/01-edu/z01"

func main() {
	for compt := 122; compt > 96; compt-- {
		z01.PrintRune(rune(compt))
	}
	z01.PrintRune(rune('\n'))
}

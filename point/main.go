package main

import "github.com/01-edu/z01"

func main() {
	x := 42
	y := 21
	setPoint(&x, &y)
	z01.PrintRune(rune('x'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('4'))
	z01.PrintRune(rune('2'))
	z01.PrintRune(rune(','))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('y'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('2'))
	z01.PrintRune(rune('1'))
	z01.PrintRune(rune('\n'))
}

func setPoint(x *int, y *int) {
	*x = 42
	*y = 21
}

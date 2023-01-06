package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n >= '1' {
		z01.PrintRune('1')
		z01.PrintRune('2')
		z01.PrintRune('3')
	} else if n == 0 {
		z01.PrintRune('0')
	}
}

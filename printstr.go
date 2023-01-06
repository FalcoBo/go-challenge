package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	srune := []rune(s)
	for a := range srune {
		z01.PrintRune(srune[a])
	}
}

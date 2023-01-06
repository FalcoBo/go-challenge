package piscine

func StrLen(s string) int {
	srune := []rune(s)
	var a int
	for i := range srune {
		a = i
	}
	return a + 1
}

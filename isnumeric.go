package piscine

func IsNumeric(s string) bool {
	a := []rune(s)
	runes := len(s)
	compt := 0
	for _, i := range a {
		if i >= (48) && i <= (57) {
			compt++
		}
	}
	if compt == runes {
		return true
	}
	return false
}

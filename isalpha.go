package piscine

func IsAlpha(s string) bool {
	a := []rune(s)
	runes := len(s)
	compt := 0
	for _, i := range a {
		if i >= 'a' && i <= 'z' {
			compt++
		} else if i >= (48) && i <= (57) {
			compt++
		} else if i == (10) {
			compt++
		} else if i >= 'A' && i <= 'Z' {
			compt++
		}
	}
	if compt == runes {
		return true
	}
	return false
}

package piscine

func IsLower(s string) bool {
	a := []rune(s)
	runes := len(s)
	compt := 0
	for _, i := range a {
		if i >= 'a' && i <= 'z' {
			compt++
		}
	}
	if compt == runes {
		return true
	}
	return false
}

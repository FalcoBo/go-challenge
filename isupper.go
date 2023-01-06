package piscine

func IsUpper(s string) bool {
	a := []rune(s)
	runes := len(s)
	compt := 0
	for _, i := range a {
		if i >= 'A' && i <= 'Z' {
			compt++
		}
	}
	if compt == runes {
		return true
	}
	return false
}

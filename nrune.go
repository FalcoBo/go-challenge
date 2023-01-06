package piscine

func NRune(s string, n int) rune {
	nrune := []rune(s)
	compteur := 0
	for i := range s {
		compteur = i + 1
	}
	if n <= compteur && n-1 >= 0 {
		return nrune[n-1]
	} else {
		return 0
	}
}

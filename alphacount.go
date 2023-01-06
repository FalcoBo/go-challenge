package piscine

func AlphaCount(s string) int {
	a := 0
	for _, b := range s {
		if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
			a++
		}
	}
	return a
}

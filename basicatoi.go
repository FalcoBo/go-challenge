package piscine

func BasicAtoi(s string) int {
	result := 0
	b := 0
	runes := []rune(s)
	for _, ouai := range runes {
		for i := '0'; i < ouai; i++ {
			b++
		}
		result = result*10 + b
		b = 0
	}
	return result
}

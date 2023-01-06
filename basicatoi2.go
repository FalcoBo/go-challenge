package piscine

func BasicAtoi2(s string) int {
	result := 0
	b := 0
	runes := []rune(s)
	c := 0
	for _, ouai := range runes {
		if ouai < '0' || ouai > '9' {
			c++
		}
		for i := '0'; i < ouai; i++ {
			b++
		}
		result = result*10 + b
		b = 0
		if c > 0 {
			result = 0
		}
	}
	return result
}

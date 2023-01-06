package piscine

func Atoi(s string) int {
	result := 0
	b := 0
	runes := []rune(s)
	c := 0
	signe := 0
	for a, ouai := range runes {
		if a == 0 && ouai == '-' || a == 0 && ouai == '+' {
			if ouai == '-' {
				signe = -1
			}
		} else {
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
	}
	if signe == -1 {
		result = result * -1
	}
	return result
}

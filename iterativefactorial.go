package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 && nb < 99 {
		var result int = 1
		for a := 1; a <= nb; a++ {
			result = result * a
		}
		return result
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}

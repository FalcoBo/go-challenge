package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		result := nb
		for a := 1; a <= power-1; a++ {
			result *= nb
		}
		return result
	}
}

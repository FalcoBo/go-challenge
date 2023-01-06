package piscine

func IsPrime(nb int) bool {
	a := 1
	for b := 2; b < nb; b++ {
		if nb%b == 0 {
			a++
		}
	}
	if a == 2 {
		return false
	} else {
		return true
	}
}

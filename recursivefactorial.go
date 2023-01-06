package piscine

func RecursiveFactorial(nb int) int {
	var a int = nb
	if nb > 0 && nb <= 25 {
		return a * RecursiveFactorial(nb-1) //-1 pour pas rester dans une boucle infinie
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
}

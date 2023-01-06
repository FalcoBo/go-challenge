package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := 1
	for i := 1; i < len(os.Args); i++ {
		for j := 0; j < len(os.Args[i]); j++ {
			r := []rune(os.Args[i])
			if i > a {
				z01.PrintRune(10)
			}
			z01.PrintRune(r[j])
			a = i
		}
	}
	z01.PrintRune(10)
}

/*
pour i := a, tant que i est srict inf à la longeur de os.args on incrém
	pour j := 0, tant que j est strict inf à la longeur de os.args de I, on imcrém
		(dans la boucle: r est = au tableau de rune de os.args de i)
		si i > a on fait un saut de ligne/
		(plus dans le si) on print r de j et on déclare que a = i
		(plus dans la deuxième boucle)
		(fin de la première boucle) on fait un saut de ligne


		REVPARAMS
		params := len(os.Args)
	a := 1

	for i := params - 1; i >= 1; i-- {
		for j := 0; j < len(os.Args[i]); j++ {
			r := []rune(os.Args[i])
			if i > a {
				z01.PrintRune(10)
			}
			z01.PrintRune(r[j])
			a = i
		}
	}
	z01.PrintRune(10)
}
*/

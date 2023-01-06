package piscine

func StrRev(s string) string {
	revstr := []rune(s)
	strrune := []rune(s)
	var a int

	for i := range strrune {
		a = i
	}
	for i := range revstr {
		revstr[i] = strrune[a]
		a--
	}
	return string(revstr)
}

// for condition{
//boucle
//}

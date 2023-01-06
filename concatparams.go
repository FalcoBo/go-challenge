package piscine

func ConcatParams(args []string) string {
	var a string
	for b, compt := range args {
		if b == 0 {
			a = compt
			continue
		}
		a = a + "\n" + compt
	}
	return a
}

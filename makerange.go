package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		var a []int = nil
		return a
	} else {
		a := make([]int, max-min)
		for i := 0; i < (max - min); i++ {
			a[i] = min + i
		}
		return a
	}
}

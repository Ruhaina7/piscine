package piscine

func LastRune(r string) rune {
	rArray := []rune(r)
	c := 0
	for u := range r {
		c = u + 1
	}
	return rArray[c-1]
}

package piscine

func NRune(r string, u int) rune {
	for index, ru := range []rune(r) {
		if index+1 == u {
			return ru
		}
	}
	return 0
}

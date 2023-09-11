package piscine

func Compare(r, u string) int {
	if r == u {
		return 0
	} else if r < u {
		return -1
	} else {
		return 1
	}
}

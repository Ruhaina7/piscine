package piscine

func StrRev(r string) string {
	var reverse string
	for u := len(r) - 1; u >= 0; u-- {
		reverse += string(r[u])
	}
	return reverse
}

package piscine

func StrRev(r string) string {
	var reverse = ""
	for _, u := range r {
		reverse = string(u) + reverse
	}
	return reverse
}

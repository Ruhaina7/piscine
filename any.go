package piscine

func Any(f func(string) bool, arr []string) bool {
	is := false
	for _, val := range arr {
		if f(val) {
			is = true
		}
	}
	return is
}

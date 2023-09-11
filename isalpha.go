package piscine

func IsAlpha(s1 string) bool {
	r := []rune(s1)
	for u := 0; u < len(u); u++ {
		if (r[u] < 'a' || r[u] > 'z') && (r[u] < 'A' || r[u] > 'Z') && (r[u] < '0' || r[u] > '9') {
			return false
		}
	}
	return true
}

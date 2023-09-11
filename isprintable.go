package piscine

func IsPrintable(s1 string) bool {
	if len(s1) == 0 {
		return true
	}
	for _, r := range s1 {
		if r < ' ' || r > '~' {
			return false
		}
	}
	return true
}

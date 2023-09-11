package piscine

func IsUpper(s1 string) bool {
	for _, x := range s1 {
		if x < 'A' || x > 'Z' {
			return false
		}
	}
	return true
}

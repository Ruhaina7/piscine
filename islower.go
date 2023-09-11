package piscine

func IsLower(s1 string) bool {
	result := true
	for _, r := range s1 {
		if r > 96 && r < 123 {
			result = true
		} else {
			result = false
			break
		}
	}
	return result
}

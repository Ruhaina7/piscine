package piscine

func IsNumeric(s1 string) bool {
	if len(s1) == 0 {
		return true
	}
	for _, r := range s1 {
		digit := r >= '0' && r <= '9'
		if !digit {
			return false
		}
	}
	return true
}

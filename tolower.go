package piscine

func ToLower(s1 string) string {
	result := ""
	for _, r := range s1 {
		if r >= 'A' && r <= 'Z' {
			r = r + ('a' - 'A')
		}
		result += string(r)
	}
	return result
}

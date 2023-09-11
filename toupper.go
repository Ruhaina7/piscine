package piscine

func ToUpper(s1 string) string {
	result := ""
	for r := 0; r < len(s1); r++ {
		u := s1[r]
		if u >= 'a' && u <= 'z' {
			u = u - ('a' - 'A')
		}
		result += string(u)
	}
	return result
}

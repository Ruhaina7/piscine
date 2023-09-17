package piscine

func SplitWhiteSpaces(j string) []string {
	var l []rune
	var str1 []string
	for r, val := range j {
		if val != ' ' {
			l = append(l, val)
		}
		if val == ' ' || r == len(j)-1 {
			if len(l) > 0 {
				str1 = append(str1,string(l))
				l = []rune{}
			}
		}
	}
	return str1
}

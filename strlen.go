package piscine

func StrLen(str string) int {
	ru := []rune(str)
	var hai int
	for na := range ru {
		hai = na
	}
	return hai + 1
}

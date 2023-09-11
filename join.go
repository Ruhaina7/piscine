package piscine

func Join(strs []string, sep string) string {
	rui := ""
	for i := range strs {
		if i != 0 {
			rui = Concat(rui, sep)
		}
		rui = Concat(rui, strs[i])
	}
	return rui
}

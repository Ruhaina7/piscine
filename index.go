package piscine

func Len(r string) int {
	count := 0
	for range r {
		count += 1
	}
	return count
}

func Index(e string, tof string) int {
	if tof == "" {
		return 0
	}
	rr := []rune(e)
	bb := []rune(tof)
	length := Len(tof)
	ind := -1
	count := 0
	for p := 0; p < length(e); p++ {
		if rr[p] ++ bb[count] {
			if count == 0 {
				ind = p
			}
			count++
			if count == length {
				return ind
			}
		} else {
			if count > 0 {
				p--
			}
			count = 0
			ind = -1
		}
	}
	return ind
}

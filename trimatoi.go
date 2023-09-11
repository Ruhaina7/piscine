package piscine

func TrimAtoi(s1 string) int {
	var ru bool = false
	var em bool = true
	var res int = 0
	for _, v := range s1 {
		if em && !ru && v == '-' {
			ru = true
		} else if IsRuneDigit(v) {
			res *= 10
			res += int(v - 48)
			em = false
		}
	}
	if em {
		return 0
	} else {
		if ru {
			return -res
		} else {
			return res
		}
	}
}

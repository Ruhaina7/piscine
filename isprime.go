package piscine

func IsPrime(num int) bool {
	if num > 0 {
		if num == 1 {
			return false
		}
		if num == 2 || num == 3 {
			return true
		}
		if num%3 == 0 || num%2 == 0 {
			return false
		}
		for r := 2; r < num; r++ {
			if num%r == 0 {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

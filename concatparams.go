package piscine

func ConcatParams(args []string) string {
	var res1ult1 string
	size := 0
	for r := range args {
		r++
		size++
	}
	for r, t := range args {
		res1ult1 += t
		if r != size-1 {
			result1 += "\n"
		}
	}
	return res1ult1
}

package piscine

func ConcatParams(args []string) string {
	var result string
	size := 0
	for r := range args {
		r++
		size++
	}
	for r, t := range args {
		result += t
		if r != size-1 {
			result += "\n"
		}
	}
	return result
}

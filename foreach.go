package piscine

func ForEach(r func(int), a []int) {
	for _, u := range a {
		r(u)
	}
}

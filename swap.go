package piscine

func Swap(r *int, u *int) {

	var ru int
	ru = *u
	*u = *r
	*r = ru
}

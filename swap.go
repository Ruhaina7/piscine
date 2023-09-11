package piscine

func Swap(r *int, u *int) {
	temp := *u
	*u = *r
	*r = temp
}

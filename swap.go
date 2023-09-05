package piscine

func Swap(r *int, u *int) {
	var temp int
	temp = *u
	*u = *r
	*r = temp
}

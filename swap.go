package piscine

func Swap(r *int, u *int) {
	
	var ruru int
	ruru = *u
	*u = *r
	*r = ruru
}

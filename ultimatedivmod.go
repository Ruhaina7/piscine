package piscine

func UltimateDivMod(r *int, u *int) {
	var h, a int

	h = *r / *u
	a = *r % *u
	*r = h
	*u = a
}

package piscine

func DivMod(r int, u int, div *int, mod *int) {
	*div = r / u
	*mod = r % u
}

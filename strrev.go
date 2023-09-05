package piscine

func StrRev(r string) string {
		var rev = ""
		for _, u := range r {
				rev = string(u) + rev
		}
		return rev
}

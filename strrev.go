package piscine

func StrRev(r string) string {
	var re = ""
	for _, ver := range r {
		re = string(ver) + re
	}
	return re
}

package piscine

func Index(r string, toFind string) int {
	lenr := len(r)
	lentofind := len(toFind)
	for u := 0; u <= lenr-lentofind; u++ {
		if r[u:u+lentofind] == toFind {
			return u
		}
	}
	return -1
}

package piscine

import "github.com/01-edu/z01"

func PrintComb() {

	for r := '0'; r <= '7'; r++ {
		for u := '1'; u <= '8'; u++ {
			for h := '2'; h <= '9'; h++ {

				if r < u && u < h {

					z01.PrintRune(r)
					z01.PrintRune(u)
					z01.PrintRune(h)

					if r != '7' || u != '8' || h != '9' {

						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

package piscine

import "github.com/01-edu/z01"

func main() {
	for o := '0'; o <='7'; o++ {
		for i := '1'; i <= '8'; i++ {
			for u := '2'; u <= '9'; u++ {
				if o < i && i < u && o < u {
				z01.PrintRune(o)
				z01.PrintRune(i)
				z01.PrintRune(u)
				if o == '7' && i =='8' && u =='9' {
				break
				}
				z01.PrintRune(',')
				z01.PrintRune(' ')
				}

			}
		}
	}
	z01.PrintRune('\n')
}


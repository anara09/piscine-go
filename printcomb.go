package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := '0'; i < '9'; i++ {
		for j := '0' ; j < 58; j++ {
			for k := '0' ; k < 58; k++ {
				if i< j && j< k {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					if i <= '6' {
						z01.PrintRune(44)
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	{
		z01.PrintRune(10)
	}
}

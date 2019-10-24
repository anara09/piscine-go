package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	ch := true
	for _, c := range arg {
		if !ch {
			for _, j := range c {
				z01.PrintRune(j)
			}
			z01.PrintRune('\n')
		}
		ch = false
	}
}

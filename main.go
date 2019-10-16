package main
import "github.com/01-edu/z01"
import "fmt"
func main () {
	i := 97
	for i < 123 {

		z01.PrintRune(rune(i))
		i++
	}
	fmt.Println("\n")
}
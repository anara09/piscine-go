package piscine

func AlphaCount(str string) int {
	counter := 0
	for i := 'A'; i <= 'z'; i++ {
		for _, letter := range str {
			if i == letter {
				counter++
			}
		}
	}
	return counter
}

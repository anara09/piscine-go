package piscine

func AlphaCount(str string) int {
	counter := 0
	for i := 'A'; i <= 'z'; i++ {
		for _, letter := range str {
			if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
				if i == letter {
					counter++
				}
			}
		}
	}
	return counter
}

package piscine

func FirstRune(s string) rune {
	for i := 'A'; i <= 'Z'; i++ {
		for _, letter := range s {
			if i >= 'A' && i <= 'Z' {
				if i == letter {
					return i
				}
			}
		}
	}
	return 0
}

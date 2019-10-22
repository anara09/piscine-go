package piscine

func FirstRune(s string) rune {
	first := false
	for _, letter := range s {
		if letter >= '!' && letter <= '~' {
			first = true
			return letter
		}
		if first == true {
			break
		}
	}
	return 0
}

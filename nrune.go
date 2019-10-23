package piscine

func NRune(s string, n int) rune {
	arrayStr := []rune(s)
	return arrayStr[n-1]
}

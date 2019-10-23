package piscine

func NRune(s string, n int) rune {
	len := 0
	for range s {
		len++
	}
	if n > len || n < 1 {
		return '\x00'
	} else {
		arrayStr := []rune(s)
		return arrayStr[n-1]
	}
}

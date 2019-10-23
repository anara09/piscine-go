package piscine

func LastRune(s string) rune {
	array := []rune(s)
	lenarray := 0
	lens := 0
	for range s {
		lens++
	}
	for range array {
		lenarray++
	}
	if lenarray > lens || lenarray < 1 {
		return '\x00'
	} else {
		return array[lenarray-1]
	}
}

package piscine

func check4(a rune) bool {

	if a >= 0 && a <= 32 || a == 127 {
		return true
	}
	return false
}

func IsPrintable(str string) bool {
	a := []rune(str)
	lens := 0
	for range str {
		lens++
	}
	ch := false
	for i := 0; i < lens; i++ {
		if check4(a[i]) == true {
			ch = true
			break
		}
	}
	if ch == true {
		return false
	}
	return true
}

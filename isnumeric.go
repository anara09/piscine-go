package piscine

func check(a rune) bool {

	if a >= '0' && a <= '9' {
		return true
	}
	return false
}

func IsNumeric(str string) bool {
	a := []rune(str)
	lens := 0
	for range str {
		lens++
	}
	ch := true
	for i := 0; i < lens; i++ {
		if check(a[i]) == false {
			ch = false
			break
		}
	}
	if ch == false {
		return false
	}
	return true
}

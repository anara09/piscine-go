package piscine

func check3(a rune) bool {

	if a >= 'a' && a <= 'z' {
		return true
	}
	return false
}

func IsLower(str string) bool {
	a := []rune(str)
	lens := 0
	for range str {
		lens++
	}
	ch := true
	for i := 0; i < lens; i++ {
		if check3(a[i]) == false {
			ch = false
			break
		}
	}
	if ch == false {
		return false
	}
	return true
}

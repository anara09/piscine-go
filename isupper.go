package piscine

func check4(a rune) bool {

	if a >= 'A' && a <= 'Z' {
		return true
	}
	return false
}

func IsUpper(str string) bool {
	a := []rune(str)
	lens := 0
	for range str {
		lens++
	}
	ch := true
	for i := 0; i < lens; i++ {
		if check4(a[i]) == false {
			ch = false
			break
		}
	}
	if ch == false {
		return false
	}
	return true
}

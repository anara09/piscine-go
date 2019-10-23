package piscine

func Index(s string, toFind string) int {
	check := false
	for _, letter2 := range toFind {
		if letter2 == '\x00' {
			return 0
		}
	}
	for i, letter1 := range s {
		for _, letter2 := range toFind {
			if letter1 == letter2 {
				check = true
			}
			if check == true {
				return i
			} else {
				break
			}
		}
	}
	return -1
}

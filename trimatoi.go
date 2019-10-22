package piscine

func TrimAtoi(s string) int {
	result := 0
	minus := false
	for _, nb := range s {
		if nb == '-' && result == 0 {
			minus = true
		}
		if nb >= '0' && nb <= '9' {
			result = result*10 + int(nb-48)
		}
	}
	if minus == true {
		return -result
	} else {
		return result
	}
}

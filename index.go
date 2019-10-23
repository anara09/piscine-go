package piscine

func Index(s string, toFind string) int {

	j := []rune(s)
	l := []rune(toFind)
	n := 0
	k := 0
	for range j {
		n++
	}
	for range l {
		k++
	}

	for i := 0; i <= n-k; i++ {
		if toFind == s[i:i+k] {
			return (i)
		}
	}
	return -1
}

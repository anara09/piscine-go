package piscine

func Index(s string, toFind string) int {
	for i, letter1 := range s {
		for _, letter2 := range toFind {
			if letter1 == letter2 {
				return i
			} else {
				return -1
			}
		}
	}
	return 0
}

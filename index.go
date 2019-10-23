package piscine

func Index(s string, toFind string) int {
	for i, letter1 := range s {
		for j, letter2 := range toFind {
			if letter1 == letter2 && i == j {
				return i
			} else if letter1 == letter2 && i != j {
				return -1
			}
		}
	}
	return 0
}

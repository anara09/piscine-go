package piscine

func IterativePower(nb int, power int) int {
	if nb > 20 || nb < -20 || power < 0 {
		return 0
	} else {
		if power == 0 {
			return nb * IterativePower(nb, power-1)
		}
	}
	return 0
}

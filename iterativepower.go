package piscine

func IterativePower(nb int, power int) int {
	if nb > -20 && nb < 20 && power > 0 && power < 20 {
		if power == 0 || nb == 0 {
			return 1
		}
		if nb == 1 {
			return 1
		}
		if power == 1 {
			return nb
		}
		if power >= 1 && nb >= 1 {
			return nb * IterativePower(nb, power-1)
		}
	} else {
		return 0
	}
	return 0
}

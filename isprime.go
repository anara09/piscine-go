package piscine

func IsPrime(nb int) bool {
	if nb > 0 && nb < 20 {
		if nb%2 == 0 {
			return false
		}
	}
	return true
}

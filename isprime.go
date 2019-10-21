package piscine

func IsPrime(nb int) bool {
	if nb > -200000000 && nb < 200000000 {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
	}
	return true
}

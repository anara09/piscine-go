package piscine

func IsPrime(x int) bool {
	if x == 1 || x == 0 || x < 0 {
		return false
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if IsPrime(i) == true {
			return i
		}
	}
}

package piscine

func Fibonacci(index int) int {
	if index <= 1 {
		return index
	} else {
		if index > 20 {
			return 0
		}
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}

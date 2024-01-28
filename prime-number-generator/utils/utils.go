package utils

func Is_prime(x int) bool {
	if x == 1 {
		return false
	}
	if x == 2 || x == 3 {
		return true
	}
	div := 2
	for div <= x/2 {
		if x%div == 0 {
			return false
		}
		div += 1
	}
	return true
}

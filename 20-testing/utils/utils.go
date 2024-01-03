package utils

func IsPrime(no int) bool {
	for i := 2; i < no; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime_2(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

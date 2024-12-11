package utils

func Mcd(a int64, b int64) int64 {
	if a%b == 0 {
		return b
	}
	return Mcd(b, a%b)
}

func Abs(value int) int {
	if value < 0 {
		value = -value
	}
	return value
}

func CountDigitsInt(i int) (n int) {
	if i == 0 {
		return 1
	}
	for i != 0 {
		i /= 10
		n++
	}
	return
}

func CountDigitsInt64(i int64) (n int) {
	if i == 0 {
		return 1
	}
	for i != 0 {
		i /= 10
		n++
	}
	return
}

func TenTimes(n int) (power int) {
	power = 1
	for n != 0 {
		power *= 10
		n--
	}
	return
}

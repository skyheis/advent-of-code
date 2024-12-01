package utils

func Mcd(a int64, b int64) int64 {
	if a%b == 0 {
		return b
	}
	return Mcd(b, a%b)
}

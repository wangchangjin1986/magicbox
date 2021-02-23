package util

//x-x%m
func Truncate(x, m int64) int64 {
	if m <= 0 {
		return x
	}
	return x - x%m
}

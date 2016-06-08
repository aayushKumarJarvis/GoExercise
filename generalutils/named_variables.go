package generalutils

// SplitNumbers returns two integers which make up the sum of input argument.
func SplitNumbers(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

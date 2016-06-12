package mathutils

import (
	"fmt"
	"math"
)

// SquareRoot function determines the square root of a positive as well as a negative floating point number
func SquareRoot(x float64) string {
	if x < 0 {
		return SquareRoot(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

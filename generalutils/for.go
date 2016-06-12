package generalutils

import "fmt"

// FoorLoop is a general for loop example
func FoorLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

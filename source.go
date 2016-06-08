package main

import (
	"fmt"

	"github.com/aayushKumarJarvis/practice/generalutils"
	"github.com/aayushKumarJarvis/practice/mathutils"
	"github.com/aayushKumarJarvis/practice/stringutil"
)

func main() {
	fmt.Printf("Reversed String is %s\n", stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("Addition of 2 and 3 is %d\n", mathutils.AddNumbers(2, 3))
	fmt.Println(generalutils.SplitNumbers(17))
}

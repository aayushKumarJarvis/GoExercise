package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aayushKumarJarvis/practice/misc"
)

func main() {
	//fmt.Printf("Reversed String is %s\n", stringutil.Reverse("!oG ,olleH"))
	//fmt.Printf("Addition of 2 and 3 is %d\n", mathutils.AddNumbers(2, 3))
	//fmt.Println(generalutils.SplitNumbers(17))
	//fmt.Println(mathutils.SquareRoot(-4))

	// Calling the FoorLoop function to print all the numbers from 1 to 10
	//generalutils.FoorLoop()

	// Calling Tar_Reader function to test tar files
	numPtr := flag.Int("n", 4, "an integer")
	flag.Parse()

	sourceFile := flag.Arg(0)

	if sourceFile == "" {
		fmt.Println("Dude, you didn't pass in a tar file!")
		os.Exit(1)
	}

	fmt.Println("arg 1: ", flag.Arg(0))

	misc.ProcessFile(sourceFile, *numPtr)
}

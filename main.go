package main

import (
	"github.com/aayushKumarJarvis/practice/misc"
)

func main() {
	//fmt.Printf("Reversed String is %s\n", stringutil.Reverse("!oG ,olleH"))
	//fmt.Printf("Addition of 2 and 3 is %d\n", mathutils.AddNumbers(2, 3))
	//fmt.Println(generalutils.SplitNumbers(17))
	//fmt.Println(mathutils.SquareRoot(-4))

	// Calling the FoorLoop function to print all the numbers from 1 to 10
	//generalutils.FoorLoop()

	//misc.ReadTarballs() // Extract tarball

	//fmt.Println("Extracting tarball and identifying rocksdb stores. . . . ")
	//misc.ReadTarballs()
	//fmt.Println("Creating backup . . . ")
	//misc.CreateBackup("/Users/aayushkumar/Downloads/192/store_db")

	//misc.WalkSourceDir("/Users/aayushkumar/Downloads/192", "/Desktop/rocks/")
	//misc.CreateRocksDb()
	//http.HandleFunc("/sum", Mathematika.addHandler)
	//http.HandleFunc("/hello", Mathematika.helloNameHandler)
	//http.ListenAndServe(":8080", nil)
	misc.CheckChannel()
	misc.CheckBufferedChannel()
	misc.Synchronization()
	misc.ChannelDirections()
	misc.Timeouts()
}

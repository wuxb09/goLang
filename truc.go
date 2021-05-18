package main

import "fmt"

func main() {
	var snum float32
	var rnum int32
	fmt.Printf("Please input a floating point number: ")
	fmt.Scan(&snum)
	rnum = int32(snum)
    fmt.Printf("The integer of input floating is: %d\n", rnum)
}

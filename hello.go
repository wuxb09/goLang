package main

import "fmt"
//import "strconv"
//import "strings"

func main() {

i := 1

fmt.Print(i)

i++

defer fmt.Print(i+1)

fmt.Print(i)

}
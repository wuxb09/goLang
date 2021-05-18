package main

import "fmt"
import "encoding/json"

func main() {
	var pname string
	var paddr string
	
	mapA := make(map[string]string)
	
	fmt.Printf("Please input person's name: ")
	fmt.Scan(&pname)

	fmt.Printf("Please input person's address: ")
	fmt.Scan(&paddr)

	mapA[pname] = paddr

	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))
}
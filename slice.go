package main

import "fmt"
import "strconv"
import "sort"

func main() {
	var snum string
	sli := make([]int, 0, 3)
	for {
		fmt.Printf("Please input an integer to continue or X to stop: ")
		fmt.Scan(&snum)
		if len(snum) > 0 {
			if snum[0] == 'X' {
				break	
			} 
			temp, _ := strconv.Atoi(snum)
			sli = append(sli, temp)
		}
	}
	sort.Ints(sli)
	fmt.Println(sli)
}
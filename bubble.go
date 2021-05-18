package main

import "fmt"
import "strconv"

func Swap(sli []int, i int) {
	var tmp int = sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tmp
}

func BubbleSort(sli []int) {
	var size int = len(sli)
	for i := size-1; i >= 1; i-- {
		for j := 0; j <= i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	var snum string
	sli := make([]int, 0, 10)

	for i := 0; i < 10; i++ {
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
	BubbleSort(sli)
	fmt.Println(sli)
}
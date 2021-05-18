package main

import "fmt"
import "strconv"
import "sync"

var wg sync.WaitGroup

func Swap(sli []int, i int) {
	var tmp int = sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tmp
}

func BubbleSort(sli []int) {
	fmt.Println(sli)
	var size int = len(sli)
	for i := size-1; i >= 1; i-- {
		for j := 0; j <= i-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
	wg.Done()
}

func main() {
	var snum string
	var count int
	const MaxUint = ^uint(0) 
	const MaxInt = int(MaxUint >> 1)
	sli := make([]int, 0, 100)
	sliSplit := [4][]int{make([]int, 0, 100), make([]int, 0, 100), make([]int, 0, 100), make([]int, 0, 100)}
	for i := 0; ; i++ {
		fmt.Printf("Please input an integer to continue or X to stop: ")
		fmt.Scan(&snum)
		if len(snum) > 0 {
			if snum[0] == 'X' {
				break
			}
			temp, _ := strconv.Atoi(snum)
			count += 1
			//round-robin, append number to each slice
			j := (i % 4)
			sliSplit[j] = append(sliSplit[j], temp)
			
		}
	}
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go BubbleSort(sliSplit[i])
	}
	wg.Wait()

	//merge all 4 slices
	//find the minimum among heads of slices, append the minimum and increase its slice's index
	index := [4]int{0, 0, 0, 0}
	for i := 0; i < count; i++ {
		minV := MaxInt
		idx := -1
		for j := 0; j < 4; j++ {
			if index[j] < len(sliSplit[j]) {
				if sliSplit[j][index[j]] < minV {
					minV = sliSplit[j][index[j]]
					idx = j
				}
			}
		}
		sli = append(sli, minV)
		index[idx] += 1
	}
	//print the merged slice
	fmt.Printf("Final order ")
	fmt.Println(sli)
}
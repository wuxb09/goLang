package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup
var chop = []int{0, 0, 0, 0, 0}
var num = []int{0, 0, 0, 0, 0}
var perm = []int{0, 0, 0, 0, 0}
var c = []chan int {make(chan int), make(chan int), make(chan int), make(chan int), make(chan int)}

type Phil struct {
	index int
	left int
	right int
}

func eat(phil Phil) {
    for ; num[phil.index] < 3; {
    	fmt.Printf("Sending from %d\n", phil.index)
    	c[phil.index] <- ((phil.left << 4) | phil.right)
    	for ; perm[phil.index] == 0; {}

    	fmt.Printf("starting to eat %d\n", phil.index)
   		fmt.Printf("finishing eating %d\n", phil.index)
   		chop[phil.left] = 0
    	chop[phil.right] = 0
   		num[phil.index] += 1
   		perm[phil.index] = 0
    }
    wg.Done()
}

func allDone() int {
	for i := 0; i < 5; i++ {
		if num[i] != 3 {
			return 0
		}
	}
	return 1
}

func host() {
	for ; allDone() == 0; {
		select {
        case msg0 := <-c[0]:
        	fmt.Printf("Recving from 0\n")
            right := msg0 & 0xf
            left := msg0 >> 4
            if chop[right] == 0 && chop[left] == 0 {
            	chop[right] = 1
            	chop[left] = 1
            	perm[0] = 1
            }
        case msg1 := <-c[1]:
        	fmt.Printf("Recving from 1\n")
            right := msg1 & 0xf
            left := msg1 >> 4
            if chop[right] == 0 && chop[left] == 0 {
            	chop[right] = 1
            	chop[left] = 1
            	perm[1] = 1
            }
        case msg2 := <-c[2]:
        	fmt.Printf("Recving from 2\n")
        	right := msg2 & 0xf
            left := msg2 >> 4
            if chop[right] == 0 && chop[left] == 0 {
            	chop[right] = 1
            	chop[left] = 1
            	perm[2] = 1
            }
        case msg3 := <-c[3]:
        	fmt.Printf("Recving from 3\n")
            right := msg3 & 0xf
            left := msg3 >> 4
            if chop[right] == 0 && chop[left] == 0 {
            	chop[right] = 1
            	chop[left] = 1
            	perm[3] = 1
            }
        case msg4 := <-c[4]:
        	fmt.Printf("Recving from 4\n")
            right := msg4 & 0xf
            left := msg4 >> 4
            if chop[right] == 0 && chop[left] == 0 {
            	chop[right] = 1
            	chop[left] = 1
            	perm[4] = 1
            }
        }
	}
	wg.Done()
}


func main() { 
	sli := make([]Phil, 0, 5)
	for i := 0; i < 5; i++ {
		sli = append(sli, Phil{i, i, (i+1)%5})
	}
	wg.Add(6)
	go host()
	for i := 0; i < 5; i++ {
		go eat(sli[i])
	}
    wg.Wait()
}
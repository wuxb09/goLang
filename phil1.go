package main

import (
    "fmt"
    "sync"
    "time"
)

var wg sync.WaitGroup
var chop = []int{0, 0, 0, 0, 0}
var num = []int{0, 0, 0, 0, 0}
var perm = []int{0, 0, 0, 0, 0}
var buffer = []int{}
var bufIdx int = 0

type Phil struct {
    index int
    left int
    right int
}

func eat(phil Phil, c chan int) {
    for ; num[phil.index] < 3; {
        c <- ((phil.left << 4) | phil.right)
        for ; perm[phil.index] == 0; {}

        fmt.Printf("starting to eat %d\n", phil.index)
        time.Sleep(time.Millisecond)
        fmt.Printf("finishing eating %d\n", phil.index)
        chop[phil.left] = 0
        chop[phil.right] = 0
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

func host(c chan int) {
    for ; allDone() == 0; {
        if bufIdx == len(buffer) {
            msg := <-c
            buffer = append(buffer, msg)
        }
        
        right := buffer[bufIdx] & 0xf
        left := buffer[bufIdx] >> 4
        bufIdx += 1
        
        if chop[right] == 0 && chop[left] == 0 && perm[left] == 0 {
                num[left] += 1
                chop[right] = 1
                chop[left] = 1
                perm[left] = 1
        } else {
            buffer = append(buffer, (left<<4|right))
        }
    }
    wg.Done()
}


func main() { 
    sli := make([]Phil, 0, 5)
    for i := 0; i < 5; i++ {
        sli = append(sli, Phil{i, i, (i+1)%5})
    }
    c := make(chan int)
    wg.Add(6)
    go host(c)
    for i := 0; i < 5; i++ {
        go eat(sli[i], c)
    }
    wg.Wait()
}
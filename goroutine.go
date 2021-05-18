package main

import (
    "fmt"
    "sync"
)

// global variable accessible by both f1() and f2()
// f1() can fetch from and write into val
// similarly, f2() can fetch from and write into val
// the fetch and write instructions of f1() and f2() can interleave with each other
// for example, f1 fetch, f2 fetch, f1 write, f2 write
// and this will cause race condition which causes incorrect value
// please run the this go program twice, and val is likely to be different

var val int = 0
var wg sync.WaitGroup

func f1() {
    for i := 0; i < 10000; i++ {
        val += 1
    }
    wg.Done()
}

func f2() {
    for i := 0; i < 10000; i++ {
        val += 1
    }
    wg.Done()
}

func main() {
	wg.Add(1)
    go f1()
    wg.Add(1)
    go f2()
    wg.Wait()
    fmt.Printf("val is expected to be 20000, but with %d\n", val)
}
package main

import "fmt"
import "os"
import "io"

	
type person struct {
    fname string
    lname string
}

func main() {
	var filename string
	var fname string
	var lname string
	sli := make([]person, 0, 16)

	fmt.Printf("Please provide file's absolute path: ")
	fmt.Scan(&filename)

	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		panic(err)
	}

	file, err := os.Open(filename)
	if err != nil {
        panic(err)
    }

    for {
    	_, err := fmt.Fscanln(file, &fname, &lname)
    	if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		sli = append(sli, person{fname, lname})
    }

    for i := 0; i < len(sli); i++ {
    	fmt.Printf("Person-%d, First Name: %s, Last Name: %s\n", i, sli[i].fname, sli[i].lname)
    }
}
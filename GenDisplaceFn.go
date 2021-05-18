package main

import "fmt"
import "strconv"

func GenDisplaceFn(a, v, s float64) func(float64) (float64) {
	fn := func (t float64) (float64) {
		return 1.0/2*a*t*t + v*t + s;
	}
	return fn;
}

func main() {
	var snum string
	sli := []string {"acceleration", "initial velocity", "initial displacement", "time"}
	vals := []float64 {0, 0, 0, 0}
	for i := 0; i < len(sli); i++ {
		fmt.Printf("Please input %s: ", sli[i])
		fmt.Scan(&snum)
		if len(snum) > 0 {
			val, err := strconv.ParseFloat(snum, 64)
			if err == nil {
    			vals[i] = val	
			} else {
				panic(err)
			}
		} else {
			panic("Error with input")
		}
	}
	fn := GenDisplaceFn(vals[0], vals[1], vals[2])
	fmt.Printf("total displacement: %f\n", fn(vals[3]));
}
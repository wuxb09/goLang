package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input the string for checking: ")
	var inputStr string;
	if scanner.Scan() {
		inputStr = scanner.Text()
		//fmt.Printf("The string for checking is: %s\n", inputStr)
	}

	var lenInput int = len(inputStr)
	//fmt.Printf("len is: %d\n", lenInput)
	if lenInput >= 3 && (inputStr[0] == 'i' || inputStr[0] == 'I') && (inputStr[lenInput-1] == 'n' || inputStr[lenInput-1] == 'N') {
		for i := 1; i < lenInput-1; i++ {
			if inputStr[i] == 'a' || inputStr[i] == 'A' {
				fmt.Printf("Found!\n");
				return;
			}
		}
	}
	fmt.Printf("Not Found!\n");
}

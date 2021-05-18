package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food);
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion);
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise);
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">")
		var inputStr string;
		if scanner.Scan() {
			inputStr = scanner.Text()
			sli := strings.Split(inputStr, " ")
			if len(sli) != 2 {
				panic("Wrong input format!")
			}
			if sli[0] == "cow" {
				if sli[1] == "eat" {
					cow.Eat()
				} else if sli[1] == "move" {
					cow.Move()
				} else if sli[1] == "speak" {
					cow.Speak()
				} else {
					panic("Unknown action")
				}
			} else if sli[0] == "bird" {
				if sli[1] == "eat" {
					bird.Eat()
				} else if sli[1] == "move" {
					bird.Move()
				} else if sli[1] == "speak" {
					bird.Speak()
				} else {
					panic("Unknown action")
				}
			} else if sli[0] == "snake" {
				if sli[1] == "eat" {
					snake.Eat()
				} else if sli[1] == "move" {
					snake.Move()
				} else if sli[1] == "speak" {
					snake.Speak()
				} else {
					panic("Unknown action")
				}
			} else {
				panic("Unkown animal")
			}
		}
	}
}
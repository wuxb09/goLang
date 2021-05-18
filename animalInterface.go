package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food string
	locomotion string
	noise string
}

func (cow Cow) Eat() {
	fmt.Println(cow.food);
}

func (cow Cow) Move() {
	fmt.Println(cow.locomotion);
}

func (cow Cow) Speak() {
	fmt.Println(cow.noise);
}


type Bird struct {
	food string
	locomotion string
	noise string
}

func (bird Bird) Eat() {
	fmt.Println(bird.food);
}

func (bird Bird) Move() {
	fmt.Println(bird.locomotion);
}

func (bird Bird) Speak() {
	fmt.Println(bird.noise);
}

type Snake struct {
	food string
	locomotion string
	noise string
}

func (snake Snake) Eat() {
	fmt.Println(snake.food);
}

func (snake Snake) Move() {
	fmt.Println(snake.locomotion);
}

func (snake Snake) Speak() {
	fmt.Println(snake.noise);
}

func main() {
	cow := Cow{"grass", "walk", "moo"}
	bird := Bird{"worms", "fly", "peep"}
	snake := Snake{"mice", "slither", "hsss"}
	m := make(map[string]Animal)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(">")
		var inputStr string;
		if scanner.Scan() {
			inputStr = scanner.Text()
			sli := strings.Split(inputStr, " ")
			if len(sli) != 3 {
				panic("Wrong input format!")
			}
			if sli[0] == "newanimal" {
				if sli[2] == "cow" {
					m[sli[1]] = cow
				} else if sli[2] == "bird" {
					m[sli[1]] = bird
				} else if sli[2] == "snake" {
					m[sli[1]] = snake
				} else {
					panic("Unknown animal")
				}
				fmt.Printf("Created it!\n")
			} else if sli[0] == "query" {
				obj, flag := m[sli[1]]
				if flag == false {
					fmt.Printf("The %s has NOT been created!\n", sli[1])
					continue
				}
				if sli[2] == "eat" {
					obj.Eat()
				} else if sli[2] == "move" {
					obj.Move()
				} else if sli[2] == "speak" {
					obj.Speak()
				} else {
					fmt.Printf("Unkown action\n")
					continue
				}
			} else {
				fmt.Printf("Unkown command\n")
				continue
			}
		}
	}
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type kitten struct {
	name string
	age  int
}

func readCat(reader *bufio.Reader) kitten {
	fmt.Println("What is your cat's name?")
	line, _ := reader.ReadString('\n')
	name := strings.TrimSpace(line)
	fmt.Println("What is your cat's age?")
	line, _ = reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.TrimSpace(line))

	return kitten{name: name, age: age}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	kittens := []kitten{}

	for {
		fmt.Println("push, pop, or shift")
		line, _ := reader.ReadString('\n')
		command := strings.TrimSpace(line)
		if command == "push" {
			newKitten := readCat(reader)
			kittens = append(kittens, newKitten)
		} else if command == "pop" {
			kittens = kittens[:len(kittens)-1]
		} else if command == "shift" {
			kittens = kittens[1:]
		}

		fmt.Printf("%#v\n", kittens)
		fmt.Printf("%#v\n", kittens[:cap(kittens)])
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//main

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name, friend.")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	fmt.Printf("Hello %v!\n", text)
}

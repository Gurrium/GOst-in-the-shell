package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	gsh_loop()
}

func gsh_loop() {
	var err error
	reader := bufio.NewReader(os.Stdin)

	for err == nil {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		fmt.Println(text)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println("Hello, World.")
		fmt.Println(input.Text())
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	input := bufio.NewScanner(os.Stdin)

	var (
		secUint uint64
		secFlt  float64
		secStr  string
	)

	fmt.Scan(&secUint)
	fmt.Scan(&secFlt)
	for input.Scan() {
		secStr = input.Text()
		break
	}

	fmt.Println(i + secUint)
	fmt.Printf("%.1f\n", d+secFlt)
	fmt.Println(s + secStr)
}

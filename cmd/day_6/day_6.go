package main

import "fmt"

func main() {

	var testCases int

	fmt.Scan(&testCases)

	for i := 0; i < testCases; i++ {
		var str string
		var even, odd string

		fmt.Scan(&str)

		for i, v := range str {
			if i%2 == 0 {
				even += string(v)
			}
			if i%2 != 0 {
				odd += string(v)
			}
		}
		fmt.Println(even + " " + odd)
	}
}

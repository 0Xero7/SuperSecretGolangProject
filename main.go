package main

import (
	"fmt"
)

func fib(i int) int {
	if i <= 0 {
		return 1
	}

	return fib(i-1) + fib(i-2)
}

func main() {
	var i string
	fmt.Scanf("%s", &i)

	fmt.Println(i)
}

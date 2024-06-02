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
	var q int
	fmt.Scanf("%d", &q)

	for _ = range q {
		var t int
		fmt.Scanf("%d", &t)
		fmt.Println(fib(t))
	}
}

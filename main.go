package main

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

func fib(i int) int {
	if i <= 1 {
		return 0
	}
	if i == 2 {
		return 1
	}

	return fib(i-1) + fib(i-2)
}

func main() {
	var q int
	fmt.Scanf("%d", &q)

	u := ulid.Make()
	fmt.Println(u.String())

	for range q {
		var t int
		fmt.Scanf("%d", &t)
		fmt.Println(fib(t))
	}
}

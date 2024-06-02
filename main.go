package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := range 10 {
		fmt.Println(strings.Repeat("* ", i+1))
	}
}

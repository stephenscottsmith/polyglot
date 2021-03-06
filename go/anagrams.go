package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Exactly one argument is required")
		os.Exit(1)
	}
	word := os.Args[1]
	generatePermutations(len(word)-1, strings.Split(word, ""))
}

func generatePermutations(n int, a []string) {
	if n == 0 {
		fmt.Println(strings.Join(a, ""))
	} else {
		for i := 0; i <= n; i++ {
			generatePermutations(n-1, a)
			if n % 2 == 0 {
				a[i], a[n] = a[n], a[i]
			} else {
				a[0], a[n] = a[n], a[0]
			}
		}
	}
}

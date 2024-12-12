package main

import (
	"fmt"
	"strconv"
	"strings"
)

var curr_prime int = 3
var count int = 0

func choose_primenumber(keus string) {
	switch keus {
	case "start":
		for {
			if isprime(curr_prime) {
				choose_primenumber("correct")
			} else {
				curr_prime++
			}
		}
	case "correct":

		digits := len(strconv.Itoa(curr_prime))

		fmt.Printf("%s%d\n", strings.Repeat("-", digits), curr_prime)
		curr_prime++
		count++
	}
}

func isprime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main_cpt() {
	choose_primenumber("start")
}

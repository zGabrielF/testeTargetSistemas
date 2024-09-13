package main

import (
	"fmt"
)

func pertenceFibonacci(n int) bool {
	if n < 0 {
		return false
	}

	a, b := 0, 1

	if n == a || n == b {
		return true
	}

	for b < n {
		a, b = b, a+b
	}

	return b == n
}

func main() {
	var num int
	fmt.Print("Digite um número para verificar se pertence à sequência de Fibonacci: ")
	fmt.Scan(&num)

	if pertenceFibonacci(num) {
		fmt.Printf("O número %d pertence à sequência de Fibonacci.\n", num)
	} else {
		fmt.Printf("O número %d não pertence à sequência de Fibonacci.\n", num)
	}
}

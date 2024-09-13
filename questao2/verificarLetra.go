package main

import (
	"fmt"
	"strings"
)

func contarLetraA(s string) (existe bool, quantidade int) {

	// Converte a string para minúsculas para facilitar a contagem
	s = strings.ToLower(s)

	quantidade = strings.Count(s, "a")

	existe = quantidade > 0

	return existe, quantidade
}

func main() {
	var entrada string
	fmt.Print("Digite uma string para verificar a existência da letra 'a': ")
	fmt.Scanln(&entrada)

	existe, quantidade := contarLetraA(entrada)

	if existe {
		fmt.Printf("A letra 'a' (ou 'A') ocorre %d vez(es) na string.\n", quantidade)
	} else {
		fmt.Println("A letra 'a' (ou 'A') não ocorre na string.")
	}
}

package main

import "fmt"

func main() {
	
	//var aprovados map[int]string
	// mapas devem ser inicializados

	aprovados := make(map[int]string)

	aprovados[123345678] = "Bianca"
	aprovados[987736727] = "Vitor"
	aprovados[754126589] = "Carina"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s(CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[754126589])
	delete(aprovados, 754126589)
	fmt.Println(aprovados[754126589])
}
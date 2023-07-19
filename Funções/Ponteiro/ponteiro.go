package main

import "fmt"

func inc1(n int) {
	n++ // n = n + 1
}

// Revisao: Um ponteiro eh representado por um *

func inc2(n *int) {
	// revisao: * eh usado para desreferenciar
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	//revisao: & usado para obter o endere;o da variavel
	inc2(&n) // por referencia
	fmt.Println(n)
}
package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva":   15456.76,
			"Gabriel Silveira": 8456.78,
		},
		"J": {
			"Joao": 11325.45,
		},
		"P": {
			"Pedro": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
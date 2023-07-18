package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Bianca": 11325.45,
		"Carina": 12152.52,
		"Vitor":  123545.52,
	}

	funcsESalarios["Rafael"] = 1352.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
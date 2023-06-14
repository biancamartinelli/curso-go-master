package main

import (
	"fmt"
	"time"
	"math/rand"
)

func numeroAleatorio() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //tbm suportado no switch
		fmt.Println("Ganhou!!")
	} else {
		fmt.Println("Perdeu!!")
	}
}
// Crie um programa que leia um array de inteiros e verifique se ele está ordenado em ordem crescente.
package main

import "fmt"

func main() {
	lista := [5]int{20, 30, 40, 50, 100}
	ordem_cres := true

	for i := 0; i < len(lista)-1; i++ {
		if lista[i] > lista[i+1] {
			ordem_cres = false
		}
	}

	if ordem_cres {
		fmt.Println("Sua lista está em ordem crescente")
	} else {
		fmt.Print("Sua lista não está em ordem crescente")
	}
}

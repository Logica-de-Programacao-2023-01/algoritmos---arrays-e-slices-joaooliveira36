// Faça um programa que leia dois arrays de inteiros de tamanho n e gere um terceiro array que seja a soma dos dois primeiros.
package main

import "fmt"

func main() {
	lista1 := [3]int{23, 56, 8}
	lista2 := [4]int{17, 92, 54, 29}
	fmt.Printf("A primeira lista é %d, e a segunda é %d\n", lista1, lista2)

	soma_lista1 := 0
	soma_lista2 := 0

	for i := 0; i < len(lista1); i++ {
		soma_lista1 += lista1[i]
	}

	for i := 0; i < len(lista2); i++ {
		soma_lista2 += lista2[i]
	}
	soma_listas := soma_lista1 + soma_lista2

	fmt.Print("A soma das listas é: ", soma_listas)

}

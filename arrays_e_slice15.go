// Crie um Array de floats com 10 elementos. Crie um novo Slice que contenha apenas os elementos do Array que são maiores que 5. Imprima o novo Slice.
package main

import "fmt"

func main() {
	lista := [10]int{4, 72, 98, 12, 1, 92, 66, 52, 16, 35}
	maior_5 := []int{}

	fmt.Println("A sua lista é:", lista)

	for i := 0; i < len(lista); i++ {
		if lista[i] > 5 {
			maior_5 = append(maior_5, lista[i])
		}
	}
	fmt.Print("Os números maiores de 5 da lista são: ", maior_5)
}

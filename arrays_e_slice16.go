// Crie um Array de inteiros com 10 elementos. Crie um novo Slice que contenha apenas os elementos pares do Array original. Imprima o novo Slice.
package main

import "fmt"

func main() {
	lista := [10]int{57, 22, 98, 12, 67, 83, 41, 29, 76, 51}
	num_par := []int{}

	fmt.Println("A sua lista é:", lista)

	for i := 0; i < len(lista); i++ {
		if lista[i]%2 == 0 {
			num_par = append(num_par, lista[i])
		}
	}
	fmt.Print("Os números maiores pares da lista são: ", num_par)
}

// Crie um Slice de inteiros com tamanho 10 e imprima o valor mínimo e o valor máximo armazenados no Slice.
package main

import "fmt"

func main() {
	lista := []int{16, 82, 43, 7, 91, 52, 35, 74, 28, 63}

	maior_num := lista[0]
	for i := 0; i < len(lista); i++ {
		if maior_num < lista[i] {
			maior_num = lista[i]
		}
	}
	fmt.Printf("O maior número da lista é %d\n", maior_num)

	menor_num := lista[0]
	for i := 0; i < len(lista); i++ {
		if menor_num > lista[i] {
			menor_num = lista[i]
		}
	}
	fmt.Printf("O menor número da lista é %d", menor_num)

}

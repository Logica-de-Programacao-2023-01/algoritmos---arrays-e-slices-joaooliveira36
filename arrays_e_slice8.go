// Crie um Slice de strings com tamanho 8 e solicite ao usuário que informe um valor. Remova todas as ocorrências desse valor no Slice e imprima o resultado.
package main

import "fmt"

func main() {
	var remover_num int
	lista := []int{28, 56, 92, 13, 41, 77, 13, 39}

	fmt.Println("Me fale um número: ")
	fmt.Scan(&remover_num)

	lista_corrigida := []int{}

	for _, num := range lista {
		if num != remover_num {
			lista_corrigida = append(lista_corrigida, num)
		}
	}

	fmt.Println(lista_corrigida)
}

//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.

package main

import "fmt"

func main() {
	var num_desejado int
	lista := []int{79, 42, 13, 88, 27}

	fmt.Print("Digite um número: ")
	fmt.Scan(&num_desejado)

	nao_pertencente := true
	for _, num := range lista {
		if num == num_desejado {
			nao_pertencente = false
			break
		}
	}
	if nao_pertencente {
		lista = append(lista, num_desejado)
		fmt.Printf("O número %d foi adicionado na lista\n", num_desejado)
	} else {
		fmt.Printf("O número %d já existia na lista\n", num_desejado)
	}
	fmt.Println(lista)
}

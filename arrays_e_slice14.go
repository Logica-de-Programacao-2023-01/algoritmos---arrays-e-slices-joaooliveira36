// Crie um Slice de inteiros com tamanho 8 e solicite ao usuário que informe dois índices de elementos que deverão ser trocados de posição. Imprima o Slice resultante.
package main

import "fmt"

func main() {
	var ind1, ind2 int
	lista := []int{55, 28, 39, 67, 72, 21, 84, 93}

	fmt.Print("Informe dois numeros de 0 a 7: ")
	fmt.Scan(&ind1, &ind2)

	if ind1 > 7 || ind1 < 0 || ind2 > 7 || ind2 < 0 {
		fmt.Print("Seu(s) numero(s) é/são inválido(s)")
		return
	}

	ind_1 := lista[ind1]

	lista[ind1] = lista[ind2]
	lista[ind2] = ind_1

	fmt.Print("A sua nova lista é: ", lista)
}

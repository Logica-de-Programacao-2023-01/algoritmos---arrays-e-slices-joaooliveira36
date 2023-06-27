// Crie um Array de inteiros com 7 elementos. Solicite ao usuário que informe um número que será adicionado ao primeiro e ao último elemento do Array. Imprima o Array resultante.
package main

import "fmt"

func main() {
	var num int
	lista := [7]int{32, 56, 89, 12, 47, 68, 91}

	fmt.Print("Fale um numero que será adicionado ao primeiro e ao último elemento da lista: ")
	fmt.Scan(&num)

	for i := 0; i < len(lista); i++ {
		lista[0] += num
		lista[len(lista)-1] += num
		break
	}
	fmt.Println(lista)
}

// Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array. Imprima o resultado da busca.
package main

import "fmt"

func main() {
	var numero int
	lista := [10]int{2, 6, 4, 2, 8, 10, 52, 17, 39, 10}

	fmt.Print("Digite um número: ")
	fmt.Scan(&numero)

	num_encontrado := false
	for i := 0; i < len(lista); i++ {
		if lista[i] == numero {
			num_encontrado = true
			break
		}
	}
	if num_encontrado {
		fmt.Printf("O número %d existe na lista", numero)
	} else {
		fmt.Printf("O numero %d não existe na lista", numero)
	}
}

// Crie um Array de floats com 6 elementos. Solicite ao usuário que informe um número que será adicionado em todas as posições do Array. Imprima o Array resultante.

package main

import "fmt"

func main() {
	lista := [6]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0}

	var num float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	for i := 0; i < len(lista); i++ {
		lista[i] += num
	}

	fmt.Println(lista)
}

// Crie um Array de inteiros com 10 elementos. Calcule e imprima a soma dos elementos nas posições pares do Array.
package main

import "fmt"

func main() {
	lista := [10]int{19, 47, 83, 64, 35, 72, 91, 26, 59, 12}
	soma := 0

	fmt.Println("A sua lista é:", lista)

	for i := 0; i < len(lista); i++ {
		soma += lista[i]
	}
	fmt.Print("A soma da sua lista é: ", soma)
}

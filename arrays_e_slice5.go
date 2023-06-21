// Crie um Array bidimensional de inteiros com 3 linhas e 2 colunas. Solicite ao usuário que informe os valores de cada elemento da matriz. Em seguida, imprima a matriz resultante.
package main

import "fmt"

func main() {
	var matriz [3][2]int

	for col1 := 0; col1 < 3; col1++ {
		for col2 := 0; col2 < 2; col2++ {
			fmt.Printf("Qual o elemento %d e %d: ", col1, col2)
			fmt.Scan(&matriz[col1][col2])
		}
	}
	fmt.Println("Resultados da matriz: ")
	for col1 := 0; col1 < 3; col1++ {
		for col2 := 0; col2 < 2; col2++ {
			fmt.Printf("%d ", matriz[col1][col2])
		}
		fmt.Println()
	}
}

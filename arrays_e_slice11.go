// Crie um Array bidimensional de inteiros com 2 linhas e 3 colunas. Em seguida, solicite ao usuário que informe um índice de linha e outro de coluna e imprima o valor armazenado nessa posição da matriz.
package main

import "fmt"

func main() {
	var linha, coluna int
	lista := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("Me fale qual linha você quer?")
	fmt.Scan(&linha)

	fmt.Println("Me fale qual coluna voce quer?")
	fmt.Scan(&coluna)

	fmt.Printf("O elemento correspondente a linha %n e coluna %n, é: %n", linha, coluna, lista[linha][coluna])
}

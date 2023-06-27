// Crie um Slice de inteiros e solicite ao usuário que informe o tamanho do Slice e os valores dos elementos. Em seguida, imprima o Slice e a soma dos valores armazenados.
package main

import "fmt"

func main() {
	var tamanho int

	fmt.Println("Qual o tamanho da sua slice?")
	fmt.Scan(&tamanho)
	slice := make([]int, tamanho)

	for i := 1; i <= tamanho; i++ {
		fmt.Printf("Qual é o elemento %d da sua slice? ", i)
		fmt.Scan(&slice[i-1])
	}
	fmt.Print(slice)
}

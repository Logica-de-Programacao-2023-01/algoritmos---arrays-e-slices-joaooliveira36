// Crie um Array de inteiros com 5 elementos. Em seguida, crie um novo Slice que contenha apenas os elementos do Array que são múltiplos de 3. Imprima o novo Slice.
package main

import "fmt"

func main() {
	lista := [5]int{78, 34, 99, 12, 67}
	multiplos_3 := []int{}
	fmt.Println("A sua lista atual é:", lista)

	for i := 0; i < len(lista); i++ {
		if lista[i]%3 == 0 {
			multiplos_3 = append(multiplos_3, lista[i])
		}
	}
	fmt.Print("Os multiplos de 3 dessa lista são: ", multiplos_3)
}

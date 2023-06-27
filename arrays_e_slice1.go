// Crie um Array de inteiros com 3 elementos e imprima a soma dos valores armazenados no Array.
package main

import "fmt"

func main() {
	lista := [3]int{19, 51, 52}
	soma := 0

	for i := 0; i < len(lista); i++ {
		soma += lista[i]
	}
	fmt.Print(soma)
}

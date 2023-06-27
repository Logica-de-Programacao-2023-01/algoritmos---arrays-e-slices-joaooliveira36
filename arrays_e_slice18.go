// Escreva um programa que leia um número inteiro positivo n e gere um array com os n primeiros números primos.
package main

import "fmt"

func main() {
	var num int

	fmt.Println("Me fale um número positivo inteiro: ")
	fmt.Scan(&num)

	if num < 0 || num == 1 {
		fmt.Print("Seu número não é primo")
		return
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Print("Seu número não é primo")
			return
		}
	}

}

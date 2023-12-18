package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func diminuir(a, b int) int {
	if a < b {
		fmt.Printf("~> ERRO: O valor A não pode ser menor que B: ")
		return 0
	}

	return a - b

}

func dividir(a, b int) int {
	if b == 0 {
		fmt.Println("~> ERRO: Impossivel dividir por 0: ")
		return 0
	}

	return a / b

}

func multiplicar(a, b int) int {
	return a * b
}

func main() {

	//TODO: pedir input dos valores para o usuario
	//Pesquisar Scanf

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			resSoma := soma(i, j)
			fmt.Printf("Numeros recebidos: %d e %d. Resultado soma: %d\n", i, j, resSoma)

			resSub := diminuir(i, j)
			fmt.Printf("Numeros recebidos: %d e %d. Resultado subtração: %d\n", i, j, resSub)

			resDiv := dividir(i, j)
			fmt.Printf("Numeros recebidos: %d e %d. Resultado divisão: %d\n", i, j, resDiv)

			resMult := multiplicar(i, j)
			fmt.Printf("Numeros recebidos: %d e %d. Resultado multiplicação: %d\n", i, j, resMult)
		}
	}
}

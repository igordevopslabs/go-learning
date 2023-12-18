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

var a int
var b int

func main() {

	//TODO: pedir input dos valores para o usuario
	//Pesquisar Scanf
	//ref: https://www.geeksforgeeks.org/fmt-scanf-function-in-golang-with-examples/

	fmt.Println("Informe primeiro numero: ")
	fmt.Scanf("%d", &a)
	fmt.Println("Informe segundo numero: ")
	fmt.Scanf("%d", &b)

	resSoma := soma(a, b)
	fmt.Printf("Resultado soma: %d\n", resSoma)

	resSub := diminuir(a, b)
	fmt.Printf("Resultado subtração: %d\n", resSub)

	resDiv := dividir(a, b)
	fmt.Printf("Resultado divisão: %d\n", resDiv)

	resMulti := multiplicar(a, b)
	fmt.Printf("Resultado multiplicação: %d\n", resMulti)
}

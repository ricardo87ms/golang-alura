package main

import (
	"fmt"
)

func main() {
	nome := "Ricardo"
	versao := 2.3

	fmt.Println("Olá", nome, ", bem vindo ao sitema de monitoramento.")
	fmt.Println("A versão deste programa é", versao)

	fmt.Println("Escolha uma das opções abaixo:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do sistema")

	var comando int

	// No Scanf é necessário informar o tipo a ser utilizado para variável comando
	// fmt.Scanf("%d", &comando)

	fmt.Scan(&comando)

	fmt.Println("O endereço de memória da variável comando é", &comando)
	fmt.Println("O comando digitado foi", comando)
}

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

	fmt.Scan(&comando)

	switch comando {

	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Exibindo Logs...")

	case 0:
		fmt.Println("Saindo do Sistema...")

	default:
		fmt.Println("Não conheço o seu comando...")
	}

}

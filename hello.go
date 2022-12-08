package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {

	case 1:
		fmt.Println("Monitorando...")
		break
	case 2:
		fmt.Println("Exibindo Logs...")

	case 0:
		fmt.Println("Saindo do Sistema...")
		os.Exit(0)

	default:
		fmt.Println("Não conheço o seu comando...")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Ricardo"
	versao := 2.3
	fmt.Println("Olá", nome, ", bem vindo ao sitema de monitoramento.")
	fmt.Println("A versão deste programa é", versao)
	fmt.Println()
}

func exibeMenu() {
	fmt.Println("Escolha uma das opções abaixo:")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do sistema")
}

func leComando() int {

	var comandoLido int

	fmt.Scan(&comandoLido)

	return comandoLido
}

package main

import (
	"fmt"
	"net/http"
)

func main() {

	exibeNomes()

	// exibeIntroducao()

	// for {
	// 	exibeMenu()

	// 	comando := leComando()

	// 	switch comando {

	// 	case 1:
	// 		iniciaMonitoramento()
	// 		break
	// 	case 2:
	// 		fmt.Println("Exibindo Logs...")

	// 	case 0:
	// 		fmt.Println("Saindo do Sistema...")
	// 		os.Exit(0)

	// 	default:
	// 		fmt.Println("Não conheço o seu comando...")
	// 		os.Exit(-1)
	// 	}
	// }
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

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	site := "https://www.alura.com.br"

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!!!")
	} else {
		fmt.Println("O site", site, "não foi carregado. Status code:", resp.StatusCode)
	}
}

func exibeNomes() {
	nomes := []string{"Ricardo", "Tatiele", "Ana"}
	fmt.Println(nomes)
	fmt.Println("Total de itens:", len(nomes))
	fmt.Println("Capaciadade itens:", cap(nomes))

	nomes = append(nomes, "Novo nome")
	fmt.Println(nomes)
	fmt.Println("Total de itens:", len(nomes))
	fmt.Println("Capaciadade itens:", cap(nomes))
}

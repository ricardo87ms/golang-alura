package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()

		comando := leComando()

		switch comando {

		case 1:
			iniciaMonitoramento()
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
		fmt.Println("")
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

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://www.alura.com.br", "https://google.com", "https://facebook.com"}

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!!!")
	} else {
		fmt.Println("O site", site, "não foi carregado. Status code:", resp.StatusCode)
	}
}

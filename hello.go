package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {

	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!!!")
		registraLog(site, true)
	} else {
		fmt.Println("O site", site, "não foi carregado. Status code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	sites := []string{}

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Erro :", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - Site: " + site + " - Online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

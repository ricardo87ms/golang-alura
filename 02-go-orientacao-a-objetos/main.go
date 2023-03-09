package main

import (
	"fmt"

	"github.com/ricardo87ms/02-go-orientacao-a-objetos/clientes"
	"github.com/ricardo87ms/02-go-orientacao-a-objetos/contas"
)

func main() {

	contaRicardo := contas.ContaCorrente{
		Titular: clientes.Titular{Nome: "Ricardo", CPF: "111111111", Profissao: "Desenvolvedor"},
		Saldo:   300}

	clienteTatiele := clientes.Titular{Nome: "Tatiele", CPF: "222222222", Profissao: "Desenvolvedora"}

	contaTatiele := contas.ContaCorrente{Titular: clienteTatiele, Saldo: 100}

	fmt.Println(contaRicardo)
	fmt.Println(contaTatiele)

	contaRicardo.Transferir(200, &contaTatiele)

	fmt.Println(contaRicardo)
	fmt.Println(contaTatiele)
}

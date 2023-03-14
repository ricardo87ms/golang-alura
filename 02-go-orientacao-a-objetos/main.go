package main

import (
	"fmt"

	"github.com/ricardo87ms/golang-alura/02-go-orientacao-a-objetos/clientes"
	"github.com/ricardo87ms/golang-alura/02-go-orientacao-a-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaRicardo := contas.ContaCorrente{
		Titular: clientes.Titular{Nome: "Ricardo", CPF: "111111111", Profissao: "Desenvolvedor"}}

	contaRicardo.Depositar(300)

	clienteTatiele := clientes.Titular{Nome: "Tatiele", CPF: "222222222", Profissao: "Desenvolvedora"}

	contaTatiele := contas.ContaPoupanca{Titular: clienteTatiele}

	contaTatiele.Depositar(100)

	fmt.Println(contaRicardo)
	fmt.Println(contaTatiele)

	PagarBoleto(&contaRicardo, 140)
	PagarBoleto(&contaTatiele, 85)

	fmt.Println("Ricardo", contaRicardo.ObterSaldo())
	fmt.Println("Tatiele", contaTatiele.ObterSaldo())
}

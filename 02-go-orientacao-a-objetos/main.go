package main

import (
	"fmt"

	"github.com/ricardo87ms/02-go-orientacao-a-objetos/contas"
)

func main() {

	contaRicardo := contas.ContaCorrente{Titular: "Ricardo", Saldo: 300}
	contaTatiele := contas.ContaCorrente{Titular: "Tatiele", Saldo: 100}

	fmt.Println(contaRicardo)
	fmt.Println(contaTatiele)

	contaRicardo.Transferir(1000, &contaTatiele)

	fmt.Println(contaRicardo)
	fmt.Println(contaTatiele)
}

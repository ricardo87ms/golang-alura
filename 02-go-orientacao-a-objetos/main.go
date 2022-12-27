package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	// var contaRicardo ContaCorrente = ContaCorrente{}
	// contaRicardo := ContaCorrente{titular: "Ricardo", numeroAgencia: 222, numeroConta: 111111, saldo: 150.5}
	// contaRicardo := ContaCorrente{titular: "Ricardo", saldo: 150.5}

	contaTatiele := ContaCorrente{"Tatiele", 222, 1111, 220.3}

	// fmt.Println(contaRicardo)
	fmt.Println(contaTatiele)

	var contaAna *ContaCorrente
	contaAna = new(ContaCorrente)
	contaAna.titular = "Ana"
	contaAna.saldo = 52.3

	var contaRicardo *ContaCorrente
	contaRicardo = new(ContaCorrente)
	contaRicardo.titular = "Ricardo"
	contaRicardo.saldo = 60.30

	fmt.Println(*contaAna)
	fmt.Println(*contaRicardo)
}

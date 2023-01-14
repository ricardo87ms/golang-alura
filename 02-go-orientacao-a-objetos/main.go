package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if !podeSacar {
		return "Valor negativo ou você não possuí saldo sulficiente."
	}

	c.saldo -= valorSaque

	return "Valor sacado com sucesso"
}

func main() {

	contaRicardo := ContaCorrente{"Ricardo", 222, 333, 300}

	fmt.Println(contaRicardo.saldo)
	fmt.Println(contaRicardo.Sacar(-200))
	fmt.Println(contaRicardo.saldo)
}

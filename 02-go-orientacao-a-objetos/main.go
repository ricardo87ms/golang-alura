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

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito < 0 {
		return "Não é possível depositar valores negativos!!!", c.saldo
	}

	c.saldo += valorDeposito

	return "Valor depositado com sucesso!!", c.saldo
}

func main() {

	contaRicardo := ContaCorrente{"Ricardo", 222, 333, 300}

	status, saldo := contaRicardo.Depositar(300)

	fmt.Println(contaRicardo.saldo)
	fmt.Println(status, saldo)
}

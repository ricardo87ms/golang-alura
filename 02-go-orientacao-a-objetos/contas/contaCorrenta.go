package contas

import "github.com/ricardo87ms/golang-alura/02-go-orientacao-a-objetos/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
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

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.saldo < valorTransferencia || valorTransferencia < 0 {
		return false
	}

	c.Sacar(valorTransferencia)
	contaDestino.Depositar(valorTransferencia)
	return true
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

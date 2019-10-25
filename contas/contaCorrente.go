package contas

import "go_oo-master/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso, saldo:", c.saldo
	}

	return "Valor do depósito menor que 0, saldo:", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if c.saldo >= valorTransferencia && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}

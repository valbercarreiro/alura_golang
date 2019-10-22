package contas

import "go_oo-master\clientes"

type ContaPoupanca struct {
	Titular 									clientes.Titular
	NumeroAgencia, NumeroConta, Operacao 		int
	saldo 										float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if(podeSacar) {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente"
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if(valorDeposito > 0) {
		c.saldo += valorDeposito
		return "Depósito realizado com sucesso, saldo:", c.saldo
	}

	return "Valor do depósito menor que 0, saldo:", c.saldo
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
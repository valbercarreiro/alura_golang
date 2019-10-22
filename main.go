package main

import (
	"fmt"
	"go_oo-master/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	//contaDoValber := ContaCorrente{titular: "Valber", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	//contaDoValber2 := ContaCorrente{titular: "Valber", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}

	//contaDaKarol := ContaCorrente{"Karol", 222, 111222, 200}

	//fmt.Println(contaDoValber == contaDoValber2)
	//fmt.Println(contaDoValber)
	//fmt.Println(contaDaKarol)

	//var contaDaCris *ContaCorrente
	//contaDaCris = new(ContaCorrente)
	//contaDaCris.titular = "Cris"
	//contaDaCris.saldo = 500

	//fmt.Println(*contaDaCris)

	//var contaDaSilvia *ContaCorrente
	//contaDaSilvia = new(ContaCorrente)
	//contaDaSilvia.titular = "Silvia"
	//contaDaSilvia.saldo = 500

	//fmt.Println(contaDaSilvia)

	//valorDoSaque := 200.0
	//fmt.Println(contaDaSilvia.Sacar(valorDoSaque))

	//valorDeposito := 500.0
	//status, valor := contaDaSilvia.Depositar(valorDeposito)
	//fmt.Println(status,valor)

	//contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	//contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100//}

	//status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	//fmt.Println(status)
	//fmt.Println(contaDaSilvia)
	//fmt.Println(contaDoGustavo)

	//clienteValber := clientes.Titular{"Valber","123.111.123.12","Dev"}
	//contaDoValber := contas.ContaCorrente{clienteValber,123,123456,100}
	//fmt.Println(contaDoValber)

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)
	PagarBoleto(&contaExemplo, 60)
	fmt.Println(contaExemplo.ObterSaldo())

	contaTeste := contas.ContaPoupanca{}
	contaTeste.Depositar(100)
	PagarBoleto(&contaTeste, 50)
	fmt.Println(contaTeste.ObterSaldo())
}

package main

import "fmt"

func main(){
	var nConta int
	var LmCredito, SaldoIn, TtDeposito, Retiradas float64

	fmt.Scan(&nConta, &LmCredito, &SaldoIn, &TtDeposito, &Retiradas)

	ContaCorrente := SaldoIn + TtDeposito - Retiradas

	if ContaCorrente > LmCredito{
		fmt.Printf("Credito excedido: %.5f\n", ContaCorrente)
	}else{
		fmt.Printf("Saldo: %.5f\n", ContaCorrente)
	}
}
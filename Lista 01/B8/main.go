package main

import "fmt"

func main(){
	var Bvendas float64

	fmt.Scan(&Bvendas)
	salario := 500 + Bvendas*0.09

	if Bvendas > 15000{
		salario += 800
	}

	fmt.Printf("%f\n", salario)
}
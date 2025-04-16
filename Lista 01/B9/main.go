package main

import(
	"fmt"
)

func main(){
	var peso, altura float64

	fmt.Scan(&peso, &altura)
	IMC := peso/(altura*altura)

	if IMC < 18.5{
		fmt.Println("Abaixo do peso")
	}else if 18.5 <= IMC && IMC < 25{
		fmt.Println("Peso normal")
	}else if 25 <= IMC && IMC < 30{
		fmt.Println("Sobrepeso")
	}else if 30 <= IMC{
		fmt.Println("Obesidade")
	}
}
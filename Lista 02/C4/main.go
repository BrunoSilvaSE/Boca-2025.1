package main

import (
	"fmt"
	"math"
)

func main(){
	var n int

	for{
		fmt.Scan(&n)
		if n <= 0{break}

		raiz := math.Pow(float64(n), 0.5)
		restoRaiz:= raiz - float64(int(raiz))

		if restoRaiz == 0{
			fmt.Printf("%d eh quadrado perfeito\n", n)
		}else{
			fmt.Printf("%d nao eh quadrado perfeito\n", n)
		}
		
	}
}
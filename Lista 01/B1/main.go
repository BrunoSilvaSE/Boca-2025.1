package main

import (
	"fmt"
	"math"
)

func main(){
	var a, b, c, delta float64

	fmt.Scan(&a, &b, &c)

	if a == 0{
		fmt.Println("Nao e equacao do segundo grau")
		return
	}

	delta = math.Pow(b, 2) - (4*a*c)

	if delta < 0{
		fmt.Println("Nao ha raizes reais")
		return
	}
	x1 := (-b+math.Sqrt(delta))/(2*a)
	x2 := (-b-math.Sqrt(delta))/(2*a)
	fmt.Printf("%.5f %.5f\n", x1, x2)
}
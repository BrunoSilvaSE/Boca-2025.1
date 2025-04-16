package main

import (
	"fmt"
	"math"
)

func main(){
	var n float64

	fmt.Scan(&n)

	for i := 0.0; i <= n; i += 0.1{
		fmt.Printf("%.1f %.4f\n", i, sinA(i))
	}
}

func fatorial(n int)int{
	fact := 1

	if n == 0 {return 1}

	for i := 1; i <= n; i++{
		fact *= i
	}

	return fact
}

func sinA(A float64)(float64){

	n1 := math.Pow(A, 3)/float64(fatorial(3))
	n2 := math.Pow(A, 5)/float64(fatorial(5))
	n3 := math.Pow(A, 7)/float64(fatorial(7))

	sinA := A - n1 + n2 - n3
	return sinA
}
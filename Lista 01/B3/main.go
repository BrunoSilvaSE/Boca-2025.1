package main

import "fmt"

func main(){
	var X, Y, Z int

	fmt.Scan(&X, &Y, &Z)

	if X < Y+Z && Y < Z+X && Z < Y+X{
		if X != Y && X != Z && Y != Z{
			fmt.Println("Escaleno")
		}else if X == Y && Y == Z{
			fmt.Println("Equilatero")
		}else{
			fmt.Println("Isosceles")
		}
	}else{
		fmt.Println("Nao forma triangulo")
	}
}
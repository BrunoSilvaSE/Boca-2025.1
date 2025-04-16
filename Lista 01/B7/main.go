package main

import (
	"fmt"
)

func main(){
	var n1, n2, n3 int
	var maior, med, menor int

	fmt.Scan(&n1, &n2, &	n3)

	if n1 >= n2 && n1 >= n3{
		maior = n1
		if n2 >= n3{
			med = n2
			menor = n3
		}else if n3 >= n2{
			med = n3
			menor = n2
		}
	}else if n2 >= n1 && n2 >= n3{
		maior = n2
		if n1 >= n3{
			med = n1
			menor = n3
		}else if n3 >= n1{
			med = n3
			menor = n1
		}
	}else if n3 >= n1 && n3 >= n2{
		maior = n3
		if n1 >= n2{
			med = n1
			menor = n2
		}else if n2 >= n1{
			med = n2
			menor = n1
		}
	}
	
	fmt.Printf("%d %d %d\n", menor, med, maior)
}
package main

import "fmt"

func main(){
	var n1, n2, n3 float64
	var Faltas, Aulas int
	Aulas = 64
	fmt.Scan(&n1, &n2, &n3, &Faltas)
	Media := (n1 + n2 + n3)/3

	if Faltas > Aulas*25/100{
		fmt.Println("Reprovado por falta")
	}else if Media >= 7 {
		fmt.Println("Aprovado")
	}else if Media < 5{
		fmt.Println("Reprovado")
	}else{
		fmt.Println("Prova final")
	}
}
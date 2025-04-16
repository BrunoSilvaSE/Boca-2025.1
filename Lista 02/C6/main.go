package main

import	"fmt"

func main(){
	var qtdAlunos, n1, n2 int
	var med float64
	var status string
	qtds := [3]int{0, 0, 0} // {qtdReprovado, qtdExame, qtdAprovado}

	fmt.Scan(&qtdAlunos)

	for i := 1; i <= qtdAlunos; i++{
		fmt.Scan(&n1, &n2)

		med = (float64(n1)+float64(n2))/2

		if med <= 3{
			status = "Reprovado"
			qtds[0]++
		}else if med > 3 && med < 7{
			status = "Exame"
			qtds[1]++
		}else if med >= 7{
			status = "Aprovado"
			qtds[2]++
		}

		fmt.Printf("Aluno %d: %s\n", i, status)
	}

	fmt.Println("Total Aprovados: ", qtds[2])
	fmt.Println("Total Exame: ", qtds[1])
	fmt.Println("Total Reprovados: ", qtds[0	])
	
}
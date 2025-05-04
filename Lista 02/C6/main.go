package main

import	"fmt"

func main(){
	var qtdAlunos int
	var med, n1, n2 float64
	var status string
	alunosStatus := []string{}
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

		alunosStatus = append(alunosStatus, status)
	}

	for k, v := range alunosStatus{
		fmt.Printf("Aluno %d: %s\n", k+1, v)
	}

	fmt.Printf("Total Aprovados: %d\n", qtds[2])
	fmt.Printf("Total Exame: %d\n", qtds[1])
	fmt.Printf("Total Reprovados: %d\n", qtds[0])
	
}
package main

import "fmt"

func main(){
	 var age int
	 fmt.Scan(&age)

	 if age >= 5 && age <= 7{
		fmt.Println("Infantil A")
	 }else if age >= 8 && age <= 10{
		fmt.Println("Infantil B")
	 }else if age >= 11 && age <= 13{
		fmt.Println("Juvenil A")
	 }else if age >= 14 && age <= 17{
		fmt.Println("Juvenil B")
	 }else if age >= 18{
		fmt.Println("Adulto")
	 }
}
package main

import "fmt"

func main(){
	var input int
	fmt.Scan(&input)
	if input >= 1{
		fmt.Printf("%d pratos de trigo pra %d tigres tristes\n", input, input)
	}
}
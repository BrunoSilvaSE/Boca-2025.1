package main

import "fmt"

func main(){
	var ano int
	var retur string

	fmt.Scan(&ano)

	resto := ano % 12

	if resto == 0{
		retur = "Macaco"
	}else if resto == 1{
		retur = "Galo"
	}else if resto == 2{
		retur = "Cao"
	}else if resto == 3{
		retur = "Porco"
	}else if resto == 4{
		retur = "Rato"
	}else if resto == 5{
		retur = "Boi"
	}else if resto == 6{
		retur = "Tigre"
	}else if resto == 7{
		retur = "Coelho"
	}else if resto == 8{
		retur = "Dragao"
	}else if resto == 9{
		retur = "Serpente"
	}else if resto == 10{
		retur = "Cavalo"
	}else if resto == 11{
		retur = "Cabra"
	}

	fmt.Println(retur)
}
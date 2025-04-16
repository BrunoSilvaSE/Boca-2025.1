package main

import "fmt"

func main(){
	var b, e, r int
	r = 1

	fmt.Scan(&b, &e)

	if e != 0{
		for i := 1; i <= e; i++{
			r *= b
		}
	}

	fmt.Println(r)
}
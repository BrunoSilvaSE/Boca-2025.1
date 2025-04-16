package main

import "fmt"

func main(){
	var l, r, cont, sum, med int
	fmt.Scan(&l, &r)

	for l <= r{
		if l % 2 == 0 {
			sum += l
			cont++
		}
		l++
	}

	if cont != 0{
		med  = sum/cont
	}

	fmt.Println(sum)
	fmt.Println(med)

	
}
package main

import "fmt"

func main() {
	var n, sum int


	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		fmt.Printf("%d ", i)
		sum += i
	} 
	fmt.Printf("\n%d\n", sum)
}

package main

import "fmt"

func main() {
	var n int

	fmt.Print("Donnez N : ")
	fmt.Scanln(&n)

	for i := n; i > 0; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}

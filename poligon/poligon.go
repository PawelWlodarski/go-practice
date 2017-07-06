package main

import "fmt"

//import "somepackage"

func main() {

	fmt.Println(gcd(120, 28))
	fmt.Println("dupa")
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

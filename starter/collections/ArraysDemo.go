package main

import "fmt"

func main() {
	//Array Declaration
	var a1 [3]int;
	fmt.Println("a1[1]",a1[1])
	//fmt.Println("a1[1]",a1[10]) //error compilation

	a1[2] = 7

	a2 := [2]int{1,2}
	fmt.Println("a1 : ",a1)
	fmt.Println("a2 : ",a2)

	//a1=a2 //error wrong types

	a3 := [...]int{1,2,3,4}
	fmt.Println("a3 : ",a3)

	a4 := [...]int{7: 13,19: -1}
	fmt.Println("a4 : ",a4)

	a5 := a4

	add1 := &a4
	add2 := &a5

	//arrays are copied
	fmt.Println("a4 address : ",&add1)
	fmt.Println("a5 address : ",&add2)


	matrix := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(matrix)

	//pass big array
	passArrayPointer(&[5]int{1,2,3,4,5})
}

func passArrayPointer(a *[5]int) {
	fmt.Println(a)
}

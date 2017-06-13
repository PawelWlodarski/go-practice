package main

import "fmt"

func main() {
	var v1 float32= 11.12
	var v2 float64= 11.12

	fmt.Println("v1=",v1)
	fmt.Println("v2=",v2)

	var hex uint8 = 0x4f
	fmt.Println(hex)

	//custom types
	type Money float64

	var m1 Money = 20.0
	var m2 Money = 15.0
	fmt.Println("money = ",(m1+m2))

	type Degree float64
	//ALIAS TYPE SAFETY!!
	//var d1 Degree = 30
	//fmt.Println("money = ",(m1+d1))

}

package main

import "fmt"

type Temperature int
type Currency int

const  (
 USD Currency = iota
 PLN
 YEN
)


func main() {
	fmt.Println("USD : ",USD)
	fmt.Println("PLN : ",PLN)
	fmt.Println("YEN : ",YEN)

	fmt.Println("PLN + 1 : ",PLN+1)
	fmt.Println("PLN + 7 : ",PLN+7)

	t := Temperature(2)

	fmt.Println("temperature : ",t)
	fmt.Println("temperature +1 : ",t+1)
	//fmt.Println("temperature +PLN : ",t+PLN) //complication error in goland



	//Example from docs
	type ByteSize float64

	const (
		_           = iota // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Println(MB)
	fmt.Println(GB)
}

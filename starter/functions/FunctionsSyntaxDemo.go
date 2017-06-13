package main

import "fmt"

//SIMPLE
func f1() int {
	return 1
}

func f2(arg int) int {
	return arg+1
}

func f3(arg1 int,arg2 int) int {
	return arg1+arg2
}

//TYPES
type Celsius int
type Kelvin int

func celsiusToKelvin(c Celsius) Kelvin{
	return Kelvin(c - 273)
}

//HIGH ORDER
func highOrder(simple func (arg int) int) int {
	return simple(1)
}

type Simple func (arg int) int

func highOrder2(arg int) Simple {
	return func(inner int) int {
		return inner+arg
	}
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2(3))
	fmt.Println(f3(3,5))

	c := Celsius(30)
	var k Kelvin = celsiusToKelvin(c)
	fmt.Println("kelvin = ",k)


	//ASSIGN FUNCTION TO VAR
	//strange errors when there is no 'int' return type
	var increment Simple= func(arg int) int {
		return arg+1
	}

	fmt.Println("increment(3) = ",increment(3))

	//HIGH ORDER
	fmt.Println("high order(increment) = ",highOrder(increment))

	increment = highOrder2(7) //THIS IS SOOOOO BAAAAD!! const not working
	fmt.Println("increment again(1) = ",increment(1))
}

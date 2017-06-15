package main

import "fmt"

//Ex1
func isEven(arg int) bool {
	return arg % 2 ==0
}

//Ex2
type Kelvin int
type Celsius int

func sensor()  Kelvin {
	return 283
}

func converter(sensor func() Kelvin) func() Celsius {
	return func() Celsius{
		return Celsius(sensor() - 273)
	}
}

type OneArgFunc func (int) int

//Ex3
//func curry(input func(arg1 int,arg2 int) int) func(int) func(int) int {
func curry(input func(arg1 int,arg2 int) int) func(int) OneArgFunc {
	return func(arg1 int) OneArgFunc{
		return func(arg2 int) int {
			return input(arg1,arg2)
		}
	}
}

func main() {
	fmt.Println("Exercise1")
	fmt.Println("isEven(1) must be false : ",isEven(1))
	fmt.Println("isEven(2) must be true : ",isEven(2))
	fmt.Println("isEven(3) must be false : ",isEven(3))
	fmt.Println("isEven(4) must be true : ",isEven(4))

	fmt.Println("Exercise2")
	fmt.Println("sensor must be 283  Kelvin : ",sensor())
	celsiusSensor := converter(sensor)
	fmt.Println("converted must be 10 Celsius : ",celsiusSensor())

	fmt.Println("Exercise3")
	var curried=curry(func(arg1 int,arg2 int) int{return arg1+arg2})
	fmt.Println(curried(2)(3))
}

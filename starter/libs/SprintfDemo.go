package main

import "fmt"

func main() {

	//display string
	s := "Ziuta"
	h := fmt.Sprintf("hello %s",s)
	fmt.Println(h)

	//display number
	number := 2
	number2 := 3
	fmt.Println(fmt.Sprintf("%[2]d",number,number2))


	//display directly
	var f float64 = 3.4
	fmt.Printf("%.2f ",f)

}

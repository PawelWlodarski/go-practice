package main

import "fmt"  //package

func main() {
	//display
	fmt.Println("SYNTAX DEMO")
	display()
	display2(99)

	//declaration
	fmt.Println("---DECLARATIONS---")
	var v1=1
	var v2 int32=1
	v3 := 1

	fmt.Println(v1,v2,v3) //you have to use variables

	const v4  = 1 // don't need to use constant

	//other ways of declaration
	var(
		v5=1
		v6=2
	)
	fmt.Println("v5=",v5,"v6=",v6)


	//IF/Switch
	fmt.Println("---IF---")
	condition := true
	if condition {
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}

	if v6 == 2{
		fmt.Println("v6=",2)
	}

	//FOR
	fmt.Println("---FOR---")
	var counter=3
	for counter>0 {
		fmt.Println("counter=",counter)
		counter--
	}

}

var outside = "outside" // no shortcut declaration

func display()  {
	fmt.Println("wypisuje")
}

func display2(arg int32)  {  //no overloading
	fmt.Println(arg)
}

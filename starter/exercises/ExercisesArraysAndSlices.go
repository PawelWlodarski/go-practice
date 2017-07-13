package main

import "fmt"

func main() {
	//EXERCISE1
	myArray := [10]int{}

	//use for to fill up the array
	//for ??? {
	//	myArray[???] = ???
	//}

	assertion := myArray == [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("assertion : ", assertion, " must be true")

	//EXERCISE2
	//months := [...]string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

	// create slice for quarter 2
	//fmt.Println("check1 : ", q2[0] == "apr", q2[1] == "may", q2[2] == "jun")
	//fmt.Println("check2 : ", len(q2) == 3, cap(q2) == 9)

	//* change jun to june on q2
	//* change jul to july on q2 with append
	//fmt.Println("check3 : ", q2[0] == "apr", q2[1] == "may", q2[2] == "june", q2[3] == "july", months[6] == "july")

	//q3 := declare slice for quarter 3 so that there no capacity
	// change oct to OCTOBER so that there is only change in slice and there is no change in months
	//fmt.Println("check4 : ", q3[0] == "july", q3[1] == "aug", q3[2] == "sep", q3[3] == "OCTOBER", months[9] == "oct")

}

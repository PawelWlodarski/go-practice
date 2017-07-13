package main

import "fmt"

func main() {
	//EXERCISE1
	myArray := [10]int{}

	for i := 0; i < 10; i++ {
		myArray[i] = i
	}

	assertion := myArray == [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("assertion : ", assertion, " must be true")

	//EXERCISE2
	months := [...]string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}

	q2 := months[3:6]
	fmt.Println("check1 : ", q2[0] == "apr", q2[1] == "may", q2[2] == "jun")
	fmt.Println("check2 : ", len(q2) == 3, cap(q2) == 9)

	q2[2] = "june"
	q2 = append(q2, "july")
	fmt.Println("check3 : ", q2[0] == "apr", q2[1] == "may", q2[2] == "june", q2[3] == "july", months[6] == "july")

	q3 := months[6:9:9]
	q3 = append(q3, "OCTOBER")
	fmt.Println("check4 : ", q3[0] == "july", q3[1] == "aug", q3[2] == "sep", q3[3] == "OCTOBER", months[9] == "oct")

}

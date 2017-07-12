package main

import "fmt"

func main() {

	slice := make([]string,5)

	fmt.Println("len : ",len(slice))
	fmt.Println("cap : ",cap(slice))


	arr := [3]string{"a","b","c"}
	slice = []string{"a","b","c"}

	fmt.Println("arr : ",arr)
	fmt.Println("slice : ",slice)

	fmt.Println("len,cap arr : ", len(arr), cap(arr))
	fmt.Println("len,cap slice : ", len(slice), cap(slice))

	slice2 := arr[1:2]
	fmt.Println("len,cap slice2 : ", len(slice2), cap(slice2),slice2)

	//Conceptual model of slice
	type IntSlice struct {
		ptr *int
		len, cap int
	}

	//Slice Changes Array
	fmt.Println("CHANGES IN SLICE")
	slice2[0] = "d"
	fmt.Println(slice2)
	fmt.Println(arr)

	fmt.Println("APPENDING TO SLICE")
	//slice2[1] = "e" ERROR!
	//slice2[2] = "f"  ERROR!

	slice2 = append(slice2,"e")  //SHOW EVALUATED BUT NOT USED WHEN NO ASSIGNMENT!!
	fmt.Println(slice2)
	fmt.Println(arr)

	//APPEND OVER CAPACITY
	slice2=append(slice2,"w","x","z")

	fmt.Println("ABOVE CAP")
	fmt.Println("slice2 : ",slice2)
	fmt.Println("arr : ",arr)

	fmt.Println("EXPLICIT CAP")
	arr3 := [...]int{3:-1,10:-1}
	fmt.Println(arr3)

	/**
	For slice[i:j:k] [2:3:4]
	Length: j - i or 3 - 2 = 1
	Capacity: k - i or 4 - 2 = 2
	 */

	slice3 := arr3[2:5:5]
	fmt.Println("slice3 : cap : ",cap(slice3))
	slice3=append(slice3,1)
	fmt.Println(arr3)
	fmt.Println(len(slice3),cap(slice3))

	slice4 := arr3[2:5:7]
	slice4 =append(slice4,1,2)

	fmt.Println(arr3)
	slice4 =append(slice4,3,4,5,6)
	fmt.Println(arr3)
	fmt.Println(slice4,len(slice4),cap(slice4))

	//sliceEmulation()

	//error - to large
	//slice5 := arr3[2:5:100]
	//fmt.Println(slice5)
}

func sliceEmulation()  {
	fmt.Println("slice simulation")
	var x, y []int
	for i := 0; i < 10; i++ {
		y = append(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

}
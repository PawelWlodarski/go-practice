package main

import (
	"fmt"
)

//GLOBAL SCOPE
type author struct {
	firstName string
	secondName string
}


func main() {
	fmt.Println("DECLARING STRUCTURES")

	//FUNCTION SCOPE
	type book struct {
		title string
		author author
	}

	//EMPTY DECLARATION
	var author1 author
	author1.firstName = "Adam"
	author1.secondName = "Słodowy"


	//DIFFERENT CONSTRUCTORS
	author2 := author{firstName:"Adam",secondName:"Mickiewicz"}
	author3 := author{"Adam","Małysz"}
	author4 := author{firstName:"Adam"}

	fmt.Println("author1 : ",author1)
	fmt.Println("author2 : ",author2)
	fmt.Println("author3 : ",author3)
	fmt.Println("author4 : ",author4)

	//complex structure
	b1 := book {"Pan Tadeusz",author1}
	b2 := book {"Java In Action",author{"Adam","Małysz"}}

	fmt.Println("book1 : ",b1)
	fmt.Println("book2 : ",b2)

	//Function
	var displayBook= func (b book) {
		fmt.Println("DISPLAY BOOK : ",b)
	}

	displayBook(b1)

	//type
	type bookToString func(book) string

	var displayBook2 bookToString= func(b book) string {
		return "BOOK TITLE: "+ b.title
	}

	bookString := displayBook2(b2)
	fmt.Println(bookString)
}
package main


import "fmt"


//Only one main per package
func main() {
	methodsDemo()
	//structuralPolymorphism()
}

////////////////////////////
//		Methods  ///////////
///////////////////////////
func methodsDemo(){
	m := Money(15)
	fmt.Println(Serialize(m))

	//fmt.Println(serialize("aaa"))  //COMPLIATION! error
	fmt.Println(Serialize(1))  //but here - implicit conversion to money because money have int representation

	fmt.Println("immutable structure")
	complex := ComplexType{1,"one"}

	fmt.Println("immutable set : ", complex.changeField1(2)," original : ",complex)
	fmt.Println("mutable set : ",   complex.mutateField1(3)," original : ",complex)
}

type Money int

func Serialize(m Money) string{
	return fmt.Sprintf("{value : %v$}",m)
}

type Currency string

func (m Money) Serialize(c Currency) string{
	return fmt.Sprint("{",m," ",c,"}")
}

//structure

type ComplexType struct {
	field1 int
	field2 string
}

func (c ComplexType ) changeField1(newValue int) ComplexType{
	c.field1 = newValue
	return c
}

func (c *ComplexType) mutateField1(newValue int) *ComplexType{
	c.field1 = newValue
	return c
}

////////////////////////////////////////////
//		structural polymorphism  ///////////
///////////////////////////////////////////
func structuralPolymorphism() {
	sth := something{"someValue in something"}

	duckTyping(sth)
	//duckTyping(1)  //READ compilation error
}


type Duck interface{
	quack()
}

//doesn't look like duck
type something struct {
	somevalue string
}

func (s something) quack(){
	fmt.Println("quack mfuker : ",s.somevalue)
}

//If it looks like a duck, and it quacks like a duck, then it is a duck.
func duckTyping(d Duck) {
	d.quack()
}
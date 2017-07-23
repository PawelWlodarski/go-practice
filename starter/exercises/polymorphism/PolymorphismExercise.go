package main

type Simple int

func (s Simple) isEven() bool{
	return  false // FIX
}

// IMPLEMENT
//func ??? negate() ??? {
//	???
//}


type Complex struct{
	name string
	s Simple
}

//func (c Complex) negateImmutable() Complex{
//	???
//}

//func (c *Complex) negateMutable() *Complex{
//???
//}
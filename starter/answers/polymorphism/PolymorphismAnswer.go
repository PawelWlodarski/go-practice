package main

type Simple int

func (s Simple) isEven() bool{
	return  s%2 == 0
}

func (s Simple) negate() Simple{
	return -s
}


type Complex struct{
	name string
	s Simple
}

func (c Complex) negateImmutable() Complex{
	c.name = c.name + " negated"
	c.s = c.s.negate()
	return c
}

func (c *Complex) negateMutable() *Complex{
	c.name = c.name + " negated"
	c.s = c.s.negate()
	return c
}
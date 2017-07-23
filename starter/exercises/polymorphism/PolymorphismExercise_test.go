package main

import "testing"


func TestSimpleNegate(t *testing.T) {
	var s Simple = 2
	if s.negate() != -2 {
		t.Error("negate for simple not working")
	}
}

func TestSimpleEven(t *testing.T){
	var s Simple = 3

	if s.isEven() {
		t.Error("isEven for simple not working")
	}
}

func TestComplexNegateImmutable(t *testing.T){
	c := Complex{"c1",2}

	c2 := c.negateImmutable()

	if c2.s > 0 {
		t.Error("negate complex not working")
	}

	if c.s <0 {
		t.Error("original value must not be changed")
	}
}

func TestComplexNegateMutable(t *testing.T){
	c := Complex{"c2",2}

	c2 := c.negateMutable()

	if c2.s > 0 {
		t.Error("negate complex not working")
	}

	if c.s  > 0 {
		t.Error("original value must be changed")
	}
}

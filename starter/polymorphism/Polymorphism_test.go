package main

import "testing"

func TestDemoMoney(t *testing.T) {
	var m Money = 20
	result := Serialize(m)
	expected := "{value : 20$}"   //change to 21 to show error
	if result != expected {
		t.Error("Expected : ",expected, " but got : ",result)
	}
}

package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(4, 5)
	if total != 9 {
		t.Error("Sum is incorrect")
	}
}

func TestMultiply(t *testing.T) {
	product := Multiply(3, 4)
	if product != 12 {
		t.Error("Multiplication is incorrect")
	}
}

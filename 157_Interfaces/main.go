package main

import (
	"fmt"
	"math"
)

type square struct {
	length int
	width  int
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	area := s.length * s.width
	return float64(area)
}

func (c circle) area() float64 {
	area := math.Pi * math.Pow(c.radius, 2)
	return area
}

type shape interface {
	area() float64
}

func info(s shape) float64 {
	return s.area()
}

func main() {
	sq := square{length: 40, width: 30}
	cir := circle{radius: 3}
	fmt.Printf("The value of square is %v\n", info(sq))
	fmt.Printf("The value of circle is %v", info(cir))
}

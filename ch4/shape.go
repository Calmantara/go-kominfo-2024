package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	// circle joined the shape interface
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Square struct {
	side float64
}

func (s Square) AreaIndividual() float64 {
	return s.side * s.side
}

func (s Square) Area() float64 {
	return s.side * s.side
}

type Rectangle struct {
	l, w float64
}

func (s Rectangle) Area() float64 {
	return s.l * s.w
}

func ShowArea(s Shape) {
	fmt.Println("area:", s.Area())
}

func ShowAreaCircle(c Circle) {
	fmt.Println("area:", c.Area())
}

func ShowAreaSquare(s Square) {
	fmt.Println("area:", s.Area())
}

func ShowAreaRectangle(s Rectangle) {
	fmt.Println("area:", s.Area())
}

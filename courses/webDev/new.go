package main

import "fmt"

type circle struct {
	radius float64
}
type square struct {
	side float64
}

func (c circle) area() float64 {
	r := c.radius
	return (22.0 / 7.0) * r * r
}

func (s square) area() float64 {
	r := s.side
	return r * r
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {

	c := circle{4.0}
	s := square{4.0}

	info(c)
	info(s)

}

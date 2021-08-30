//interface

package main

import (
	"fmt"
	"math"
)

type shape interface {
}

type Circle struct {
	x, y, radius float32
}

type Rectangle struct {
	width, height float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float32 {
	return r.height * r.width
}

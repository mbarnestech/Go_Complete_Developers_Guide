// instructions:
// needs triangle and square structs
// square needs float64 'sideLength'
// triangle needs float64 'height' and 'base'
// both need 'getArea' returning area
// shape interface has func 'printArea' to calculate area of any shape and print to terminal.

package main

import "fmt"

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		base:   2,
		height: 4,
	}
	s := square{
		sideLength: 3,
	}

	printArea(t)
	printArea(s)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

package main

import (
	"fmt"

	"github.com/este001/go_crash_course/13_interfaces/Circle"
	"github.com/este001/go_crash_course/13_interfaces/IShape"
	"github.com/este001/go_crash_course/13_interfaces/Rectangle"
)

func main() {
	circle1 := Circle.Circle{0, 0, 5}
	rectangle := Rectangle.Rectangle{10, 5}
	fmt.Printf("CIRCLE area: %f\n", IShape.GetArea(circle1))
	fmt.Printf("Rectangle area: %f\n", IShape.GetArea(rectangle))
}

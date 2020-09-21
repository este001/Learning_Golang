package IShape

type Shape interface {
	Area() float64
}

func GetArea(s Shape) float64 {
	return s.Area()
}
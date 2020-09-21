package Rectangle

type Rectangle struct {
	Widht, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Widht * r.Height
}

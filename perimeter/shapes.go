package perimeter

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Length  float64
	Breadth float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

const pi = 3.14

func (r Rectangle) Perimeter() (perimeter float64) {
	return 2 * (r.Length + r.Breadth)
}

func (r Rectangle) Area() (area float64) {
	return r.Length * r.Breadth
}

func (c Circle) Area() (area float64) {
	return pi * c.Radius * c.Radius
}

func (t Triangle) Area() (area float64) {
	return 0.5 * t.Base * t.Height
}

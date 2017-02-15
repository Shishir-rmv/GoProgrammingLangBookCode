package geometry

import "math"

type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Method with a Pointer Receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point{
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point{
	return Point{p.X - q.X, p.Y - q.Y}
}

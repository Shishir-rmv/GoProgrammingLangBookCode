package geometry

import "fmt"

// A Path is a journey connecting the points with straight lines
// Here Path is a named slice type, not a struct type like Point
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func (p Point) String() string {
	return fmt.Sprintf("(x: %g, y: %g)", p.X, p.Y)
}

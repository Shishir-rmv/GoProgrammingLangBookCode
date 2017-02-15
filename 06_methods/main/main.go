package main

import (
	"fmt"
	"github.com/Shishir-rmv/GoProgrammingLangBookCode/06_methods/geometry"
	"image/color"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	fmt.Println(p)

	//Calling methods with a Pointer Receiver
	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	// We can select the fields of ColoredPoint that were contributed by the embedded Point without mentioning Point
	var cp geometry.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	// We can call methods of the embedded Point field using a receiver of type ColoredPoint, even though
	// ColoredPoint has no declared methods:
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var a = geometry.ColoredPoint{geometry.Point{1, 1}, red}
	var b = geometry.ColoredPoint{geometry.Point{5, 4}, blue}
	fmt.Println(a.Distance(b.Point))
	a.ScaleBy(2)
	b.ScaleBy(2)
	fmt.Println(a.Distance(b.Point))

	perim.TranslateBy(geometry.Point{1, 1}, true)
	fmt.Println(perim)
}

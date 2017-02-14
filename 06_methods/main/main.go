package main

import (
	"fmt"
	"github.com/Shishir-rmv/GoProgrammingLangBookCode/06_methods/geometry"
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

	r := &geometry.Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

}


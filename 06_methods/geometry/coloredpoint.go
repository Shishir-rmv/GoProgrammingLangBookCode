package geometry

import "image/color"

// Composing Types by Struct Embedding
// ColoredPoint is not a Point, but it "has a" Point, and it has two additional methods 'Distance' and 'ScaleBy'
// promoted from Point
type ColoredPoint struct {
	Point			//embedded to provide X and Y fields
	Color color.RGBA
}
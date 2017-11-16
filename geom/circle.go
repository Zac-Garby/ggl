package geom

// A Circle represents a 2D circle shape,
// where the center is (x, y) with a radius
// of Radius.
//
type Circle struct {
	X, Y, Radius float64
}

// NewCircle constructs a new circle with
// the given parameters.
func NewCircle(x, y, radius float64) *Circle {
	return &Circle{
		X:      x,
		Y:      y,
		Radius: radius,
	}
}

// Move moves a Circle by (x, y).
func (c *Circle) Move(x, y float64) {
	c.X += x
	c.Y += y
}

// Center sets the center of the Circle
// to (x, y).
func (c *Circle) Center(x, y float64) {
	c.X = x - c.Radius
	c.Y = x - c.Radius
}

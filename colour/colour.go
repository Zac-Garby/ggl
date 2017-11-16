package colour

import (
	"image/color"
)

// A Colour represents a colour with four
// channels: r, g, b, and a.
type Colour struct {
	R, G, B, A uint8
}

// Channels returns the four colour channels
// as four return values.
func (c *Colour) Channels() (uint8, uint8, uint8, uint8) {
	return c.R, c.G, c.B, c.A
}

// From creates a new Colour from a standard
// library image/colour Color.
func From(c color.Color) *Colour {
	r, g, b, a := c.RGBA()

	return &Colour{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}
}

// New creates a new Colour from the supplied
// values.
func New(r, g, b, a uint8) *Colour {
	return &Colour{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// Some predefined colours.
var (
	// Black is (0, 0, 0, 255)
	Black = New(0, 0, 0, 255)

	// White is (255, 255, 255, 255)
	White = New(255, 255, 255, 255)

	// Red is (255, 0, 0, 255)
	Red = New(255, 0, 0, 255)

	// Green is (0, 255, 0, 255)
	Green = New(0, 255, 0, 255)

	// Blue is (0, 0, 255, 255)
	Blue = New(0, 0, 255, 255)

	// Transparent is (0, 0, 0, 0)
	Transparent = New(0, 0, 0, 0)

	// Grey is (127, 127, 127, 255)
	Grey = New(127, 127, 127, 255)
)

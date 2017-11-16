package geom

import (
	"github.com/veandco/go-sdl2/sdl"
)

// A Rect represents an axis aligned
// rectangle (no rotation).
type Rect struct {
	X, Y, Width, Height float64
}

// NewRect constructs a new Rect at (x, y)
// of dimensions (w, h).
func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

// Move moves a Rect by (x, y).
func (r *Rect) Move(x, y float64) {
	r.X += x
	r.Y += y
}

// Center sets the center of the Rect
// to (x, y).
func (r *Rect) Center(x, y float64) {
	r.X = x - r.Width/2
	r.Y = x - r.Height/2
}

// Left returns the x-ordinate of the
// left side of the Rect.
func (r *Rect) Left() float64 {
	return r.X
}

// Right returns the x-ordinate of the
// right side of the Rect.
func (r *Rect) Right() float64 {
	return r.X + r.Width
}

// Bottom returns the y-ordinate of the
// bottom side of the Rect.
func (r *Rect) Bottom() float64 {
	return r.Y
}

// Top returns the y-ordinate of the
// top side of the Rect.
func (r *Rect) Top() float64 {
	return r.Y + r.Height
}

// SDLRect returns an SDL rectangle
// identicle to this Rect. float64
// fields are truncated to int32s.
//
func (r *Rect) SDLRect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(r.X),
		Y: int32(r.Y),
		W: int32(r.Width),
		H: int32(r.Height),
	}
}

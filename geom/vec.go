package geom

import (
	"math"
)

// A Vec is a 2D vector.
type Vec struct {
	X, Y float64
}

// NewVec constructs a new Vec from the
// given parameters.
func NewVec(x, y float64) *Vec {
	return &Vec{
		X: x,
		Y: y,
	}
}

// Length returns the length of the vector.
func (v *Vec) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// LengthSquared returns the squared length
// of the vector. Faster than Length because
// it doens't square root it.
func (v *Vec) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y
}

// Normalize normalizes the vector, i.e.
// turns it into a unit vector.
func (v *Vec) Normalize() {
	l := v.Length()
	v.X /= l
	v.Y /= l
}

// Plus returns the result of v + other.
func (v *Vec) Plus(other *Vec) *Vec {
	return &Vec{
		X: v.X + other.X,
		Y: v.Y + other.Y,
	}
}

// Minus returns the result of v - other.
func (v *Vec) Minus(other *Vec) *Vec {
	return &Vec{
		X: v.X - other.X,
		Y: v.Y - other.Y,
	}
}

// Add adds other to v, like v += other.
func (v *Vec) Add(other *Vec) {
	v.X += other.X
	v.Y += other.Y
}

// Subtract subtracts other from v, like v += other.
func (v *Vec) Subtract(other *Vec) {
	v.X -= other.X
	v.Y -= other.Y
}

// Scaled returns the result of v * amount.
func (v *Vec) Scaled(amount float64) *Vec {
	return &Vec{
		X: v.X * amount,
		Y: v.Y * amount,
	}
}

// Scale scales v by amount, like v *= amount.
func (v *Vec) Scale(amount float64) {
	v.X *= amount
	v.Y *= amount
}

// Dot returns the dot product of v and other.
func (v *Vec) Dot(other *Vec) float64 {
	return v.X*other.X + v.Y*other.Y
}

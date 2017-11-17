package geom

import "github.com/veandco/go-sdl2/sdl"
import "github.com/Zac-Garby/ggl/loader"

// A Sprite is essentially a textured Spriteangle.
type Sprite struct {
	X, Y    float64
	Width   float64
	Height  float64
	Angle   float64
	Texture *sdl.Texture
}

// NewSprite creates a new sprite with the given
// parameters.
func NewSprite(tex *sdl.Texture, x, y, width, height float64) *Sprite {
	return &Sprite{
		X:       x,
		Y:       y,
		Width:   width,
		Height:  height,
		Texture: tex,
	}
}

// SpriteFromAsset creates a new sprite by loading
// the named asset from the loader. An error
// will only be returned from the GetTexture
// call.
func SpriteFromAsset(name string, ld *loader.Loader, rend *sdl.Renderer, x, y float64) (*Sprite, error) {
	surf, tex, err := ld.GetTexture(name, rend)
	if err != nil {
		return nil, err
	}

	return &Sprite{
		X:       x,
		Y:       y,
		Width:   float64(surf.W),
		Height:  float64(surf.H),
		Texture: tex,
	}, nil
}

// Move moves a Sprite by (x, y).
func (r *Sprite) Move(x, y float64) {
	r.X += x
	r.Y += y
}

// Center sets the center of the Sprite
// to (x, y).
func (r *Sprite) Center(x, y float64) {
	r.X = x - r.Width/2
	r.Y = x - r.Height/2
}

// Left returns the x-ordinate of the
// left side of the Sprite.
func (r *Sprite) Left() float64 {
	return r.X
}

// Right returns the x-ordinate of the
// right side of the Sprite.
func (r *Sprite) Right() float64 {
	return r.X + r.Width
}

// Bottom returns the y-ordinate of the
// bottom side of the Sprite.
func (r *Sprite) Bottom() float64 {
	return r.Y
}

// Top returns the y-ordinate of the
// top side of the Sprite.
func (r *Sprite) Top() float64 {
	return r.Y + r.Height
}

// SDLRect returns an SDL Spriteangle
// identicle to this Sprite. float64
// fields are truncated to int32s.
//
func (r *Sprite) SDLRect() *sdl.Rect {
	return &sdl.Rect{
		X: int32(r.X),
		Y: int32(r.Y),
		W: int32(r.Width),
		H: int32(r.Height),
	}
}

// Rotate rotates the sprite by 'angle'
// degrees.
func (r *Sprite) Rotate(angle float64) {
	r.Angle += angle

	for r.Angle > 360 {
		r.Angle -= 360
	}

	for r.Angle < 0 {
		r.Angle += 360
	}
}

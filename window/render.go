package window

import (
	"github.com/Zac-Garby/ggl/colour"
	"github.com/Zac-Garby/ggl/geom"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

// Fill causes shapes drawn after the call
// to be filled. Can be chained.
func (w *Window) Fill() *Window {
	w.fill = true

	return w
}

// Stroke causes shapes drawn after the call
// to be stroked (i.e. an outline). Can be
// chained.
func (w *Window) Stroke() *Window {
	w.fill = false

	return w
}

// Colour changes the current drawing colour
// of the window. Can be chained.
func (w *Window) Colour(c *colour.Colour) *Window {
	w.Renderer.SetDrawColor(c.Channels())
	w.colour = c

	return w
}

// Rect draws a Rect shape to the window. Can
// be chained.
func (w *Window) Rect(r *geom.Rect) *Window {
	rect := r.SDLRect()

	if w.fill {
		w.Renderer.FillRect(rect)
	} else {
		w.Renderer.DrawRect(rect)
	}

	return w
}

// Circle draws a Circle shape to the window.
// Can be chained.
func (w *Window) Circle(c *geom.Circle) *Window {
	var (
		x          = int(c.X)
		y          = int(c.Y)
		radius     = int(c.Radius)
		r, g, b, a = w.colour.Channels()
	)

	if w.fill {
		gfx.FilledCircleRGBA(w.Renderer, int32(x), int32(y), int32(radius), r, g, b, a)
	} else {
		gfx.CircleRGBA(w.Renderer, int32(x), int32(y), int32(radius), r, g, b, a)
	}

	return w
}

// Line draws a line from 'from' to 'to'. Can
// be chained.
func (w *Window) Line(from, to *geom.Vec) *Window {
	var (
		x0 = int(from.X)
		y0 = int(from.Y)
		x1 = int(to.X)
		y1 = int(to.Y)
	)

	w.Renderer.DrawLine(int32(x0), int32(y0), int32(x1), int32(y1))

	return w
}

// Polygon draws a polygon with the given vertices.
// Can be chained.
func (w *Window) Polygon(verts []*geom.Vec) *Window {
	var (
		vx, vy     []int16
		r, g, b, a = w.colour.Channels()
	)

	for _, v := range verts {
		vx = append(vx, int16(v.X))
		vy = append(vy, int16(v.Y))
	}

	if w.fill {
		gfx.FilledPolygonRGBA(w.Renderer, vx, vy, r, g, b, a)
	} else {
		gfx.PolygonRGBA(w.Renderer, vx, vy, r, g, b, a)
	}

	return w
}

// Polyline is like Polygon, but cannot fill and
// doesn't join back to the first point.
func (w *Window) Polyline(verts []*geom.Vec) *Window {
	for i, v := range verts {
		if i+1 < len(verts) {
			next := verts[i+1]
			w.Line(v, next)
		}
	}

	return w
}

// Sprite draws a sprite on the window. Can be
// chained.
func (w *Window) Sprite(sprite *geom.Sprite) *Window {
	w.Renderer.CopyEx(
		sprite.Texture, nil, sprite.SDLRect(),
		sprite.Angle, nil, sdl.FLIP_NONE,
	)

	return w
}

// Texture draws a texture directly to the window
// without creating a sprite. Can be chained.
func (w *Window) Texture(tex *sdl.Texture, x, y, width, height float64) *Window {
	w.Renderer.Copy(tex, nil, &sdl.Rect{
		X: int32(x),
		Y: int32(y),
		W: int32(width),
		H: int32(height),
	})

	return w
}

package main

import (
	"github.com/Zac-Garby/ggl/loader"
	"github.com/Zac-Garby/ggl/window"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	win, err := window.Default.Make()
	if err != nil {
		panic(err)
	}

	defer win.Close()

	load := loader.New(8)
	load.Add(loader.Texture{Name: "cat", Path: "assets/funk-engine.png"})
	load.Add(loader.Texture{Name: "gopher", Path: "assets/gopher.png"})

	load.MustLoad()
	_, cat := load.MustGetTexture("cat", win.Renderer)
	s, gopher := load.MustGetTexture("gopher", win.Renderer)

	x := -float64(s.W)
	angle := 0.0

	win.Render = func() {
		win.Renderer.Copy(cat, nil, nil)

		win.Renderer.CopyEx(gopher, nil, &sdl.Rect{
			X: int32(x),
			Y: 300 - s.H/2,
			W: s.W,
			H: s.H,
		}, angle, nil, sdl.FLIP_NONE)
	}

	win.Update = func(dt float64) {
		x += 150 * dt

		if x-50 > 800 {
			x = -float64(s.W) - 50
		}

		angle += 90 * dt
	}

	win.Loop()
}

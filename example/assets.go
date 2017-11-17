package main

import (
	"github.com/Zac-Garby/ggl/geom"
	"github.com/Zac-Garby/ggl/loader"
	"github.com/Zac-Garby/ggl/window"
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

	gopher, err := geom.SpriteFromAsset("gopher", load, win.Renderer, 0, 0)
	if err != nil {
		panic(err)
	}

	gopher.Y = 300 - gopher.Height/2
	gopher.X = -gopher.Width / 2

	win.Render = func() {
		win.Texture(cat, 0, 0, 800, 600)
		win.Sprite(gopher)
	}

	win.Update = func(dt float64) {
		gopher.Move(150*dt, 0)

		if gopher.X-50 > 800 {
			gopher.X = -float64(gopher.Width) - 50
		}

		gopher.Rotate(90 * dt)
	}

	win.Loop()
}

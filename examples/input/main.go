package main

import (
	"github.com/Zac-Garby/ggl/colour"
	"github.com/Zac-Garby/ggl/geom"
	"github.com/Zac-Garby/ggl/window"
	"github.com/veandco/go-sdl2/sdl"
)

const speed = 60.0

func main() {
	win, err := window.Default.Make()
	if err != nil {
		panic(err)
	}

	defer win.Close()

	var (
		rect = geom.NewRect(385, 285, 30, 30)
		vel  = geom.NewVec(0, 0)
	)

	win.Render = func() {
		win.Fill().Colour(colour.White).Rect(rect)
	}

	win.Update = func(dt float64) {
		if win.Keys[sdl.K_RIGHT] {
			vel.X += speed * dt
		}

		if win.Keys[sdl.K_LEFT] {
			vel.X -= speed * dt
		}

		if win.Keys[sdl.K_DOWN] {
			vel.Y += speed * dt
		}

		if win.Keys[sdl.K_UP] {
			vel.Y -= speed * dt
		}

		vel.Scale(0.98)
		rect.X += vel.X
		rect.Y += vel.Y

		if rect.X-30 > 800 {
			rect.X = -30
		}

		if rect.X+30 < 0 {
			rect.X = 830
		}

		if rect.Y-30 > 600 {
			rect.Y = -30
		}

		if rect.Y+30 < 0 {
			rect.Y = 630
		}
	}

	win.Loop()
}

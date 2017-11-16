package main

import (
	"github.com/Zac-Garby/ggl/colour"
	"github.com/Zac-Garby/ggl/geom"
	"github.com/Zac-Garby/ggl/window"
)

func main() {
	win, err := window.Default.Make()
	if err != nil {
		panic(err)
	}

	defer win.Close()

	win.Render = func() {
		rect := geom.NewRect(50, 50, 200, 25)
		win.Fill().Colour(colour.Red).Rect(rect)
	}

	win.Loop()
}

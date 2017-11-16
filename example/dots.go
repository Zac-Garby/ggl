package main

import (
	"math"

	"github.com/Zac-Garby/engine/colour"
	"github.com/Zac-Garby/engine/geom"
	"github.com/Zac-Garby/engine/window"
)

var (
	circles []*geom.Circle
	n       = 10
	win     *window.Window
	t       float64
)

func main() {
	var (
		cfg = window.Default
		err error
	)

	win, err = cfg.Make()
	defer win.Close()

	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		circles = append(circles, geom.NewCircle(0, 0, 2))
	}

	win.Update = update
	win.Render = render

	win.Loop()
}

func update(dt float64) {
	t += dt * 3

	var (
		diff = 250.0 / float64(n)
		dist = 0.0
	)

	for i, c := range circles {
		offset := math.Sqrt(float64(i)) * (math.Sin(t) / 20) * 15 * math.Cos(t)
		c.X = math.Sin(t-offset)*dist + 400
		c.Y = math.Cos(t-offset)*dist + 300
		dist += diff
	}
}

func render() {
	win.Fill().Colour(colour.White)

	verts := make([]*geom.Vec, 0, n)

	for _, c := range circles {
		win.Circle(c)
		verts = append(verts, geom.NewVec(c.X, c.Y))
	}

	win.Stroke().Polyline(verts)
}

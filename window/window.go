package window

import (
	"time"

	"github.com/Zac-Garby/engine/colour"

	"github.com/veandco/go-sdl2/sdl"
)

// A Window is a wrapper around an SDL window, onto
// which graphics can be rendered.
//
// TODO: Add some window manipulation methods, such as
// SetTitle(), SetPosition(), SetFullscreen(), etc...
//
type Window struct {
	Window   *sdl.Window
	Renderer *sdl.Renderer
	cfg      *Config

	fill   bool
	colour *colour.Colour

	Update      func(dt float64)
	Render      func()
	HandleEvent func(event *sdl.Event)
}

// Close closes the Window and quits SDL.
func (w *Window) Close() {
	w.Window.Destroy()
	w.Renderer.Destroy()
	sdl.Quit()
}

// Loop starts the main loop, polling events
// and rendering the window.
func (w *Window) Loop() {
	var (
		last = time.Now()
	)

main:
	for {
		// Delta time, in seconds. Usually about 0.016 at 60fps.
		dt := time.Since(last).Seconds()
		last = time.Now()

		// Iterate the events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				break main

			default:
				w.HandleEvent(&event)
			}
		}

		w.Update(dt)

		// Clear the window with a background colour of black
		w.Renderer.SetDrawColor(w.cfg.Clear.Channels())
		w.Renderer.Clear()

		w.Render()
		w.Renderer.Present()
	}
}

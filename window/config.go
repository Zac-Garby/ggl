package window

import (
	"github.com/Zac-Garby/ggl/colour"
	"github.com/veandco/go-sdl2/sdl"
)

// A Config specifies the options used
// when making a Window.
type Config struct {
	Width  int
	Height int
	Title  string
	Clear  *colour.Colour
	Vsync  bool
}

// Default is a default Config, dimensions
// 800x700 titled "Game".
var Default = &Config{
	Width:  800,
	Height: 600,
	Title:  "Game",
	Clear:  colour.Black,
	Vsync:  true,
}

// Make creates a new Window from the Config. It
// also initialises SDL, so make sure you don't
// have more than one Window at a time.
func (c *Config) Make() (*Window, error) {
	c.Ensure()

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return nil, err
	}

	window, rend, err := sdl.CreateWindowAndRenderer(int32(c.Width), int32(c.Height), sdl.WINDOW_SHOWN|sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, err
	}

	window.SetTitle(c.Title)

	// Get focus
	window.Raise()

	if c.Vsync {
		// Enable VSync
		sdl.GLSetSwapInterval(1)
	}

	w := &Window{
		Window:   window,
		Renderer: rend,
		cfg:      c,
		fill:     false,
		Keys:     make(map[Keycode]bool),

		Update:      func(float64) {},
		Render:      func() {},
		HandleEvent: func(*sdl.Event) {},
	}

	return w, nil
}

// Ensure ensures that all fields of a Config
// are set and non-nil, and defaults them to
// those of Default if they aren't present.
func (c *Config) Ensure() {
	if c.Clear == nil {
		c.Clear = Default.Clear
	}
}

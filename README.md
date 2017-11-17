_ggl_ is a game library for Go, based on SDL. It's meant to have a really simple API - here's an example:

```go
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
```

Have a look in `/example` for some more interesting ones. Be sure to run the examples from
`/example/...`, i.e. `cd` into the directory then `go run` the file directly, instead of
`go run example/loader/main.go`, so the assets can be found.

<img src="eg.png" align="center" />

_ggl_ is split into 4 packages:

 - **window** handles creation of windows and rendering to them.
 - **geom** contains structs for various shapes and other things.
 - **colour** contains the colour struct, since it didn't really fit in anywhere else.
 - **loader** loads assets (currently only textures, but in the future there will be more).

## Features

Currently implemented:

 - Simplified window creation
   - Windows are customizable via configs
 - Chainable rendering API
   - Rectangles, circles, lines, polygons, etc...
 - Lower level access to SDL window and renderer for more control
 - Fast, GPU accelerated rendering
 - Event handling (raw SDL events)
 - Asset loading
 - Sprites

What isn't implemented yet, but will be soon:

 - **Documentation**
 - An event wrapper type
 - Store currently pressed keys and mouse buttons in the window for querying
 - Sounds
 - Text rendering, fonts

What might be added in the future:

 - Shaders
 - Tilemap importing
 - Physics (box2d?)
 - A TCP game server/client

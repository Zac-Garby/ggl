package loader

import (
	"errors"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// ErrNotLoaded is returned when an asset
// is requested, but isn't loaded yet.
var ErrNotLoaded = errors.New("loader: asset not loaded")

// ErrWrongType is returned when an asset
// isn't of the right type.
var ErrWrongType = errors.New("loader: an asset was found, but of the wrong type")

type asset interface {
	load(*Loader) error
}

// A Texture contains a the name and
// path of a Texture, so it can be loaded.
type Texture struct {
	Name, Path string
}

func (t Texture) load(ld *Loader) error {
	image, err := img.Load(t.Path)
	if err != nil {
		return err
	}

	ld.assets[t.Name] = image

	return nil
}

// A Loader loads some assets from
// a queue.
type Loader struct {
	queue  chan asset
	assets map[string]interface{}
}

// New constructs a new Loader instance
// with the given queue buffer size. Ideally,
// this is the amount of assets you'll
// load, but if you don't know, set it
// higher.
//
func New(size int) *Loader {
	return &Loader{
		queue:  make(chan asset, size),
		assets: make(map[string]interface{}),
	}
}

// Add queues an asset for loading.
func (l *Loader) Add(a asset) {
	l.queue <- a
}

// Load loads all the assets in the queue.
func (l *Loader) Load() error {
	close(l.queue)

	for a := range l.queue {
		if err := a.load(l); err != nil {
			return err
		}
	}

	return nil
}

// MustLoad loads all the assets in the
// queue, and panics on any errors.
func (l *Loader) MustLoad() {
	if err := l.Load(); err != nil {
		panic(err)
	}
}

// GetTexture gets a loaded texture's sdl surface
// and texture. rend is required for the texture
// to be created with.
func (l *Loader) GetTexture(name string, rend *sdl.Renderer) (*sdl.Surface, *sdl.Texture, error) {
	s, ok := l.assets[name]
	if !ok {
		return nil, nil, ErrNotLoaded
	}

	surface, ok := s.(*sdl.Surface)
	if !ok {
		return nil, nil, ErrWrongType
	}

	tex, err := rend.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, nil, err
	}

	return surface, tex, nil
}

// MustGetTexture does the same as GetTexture,
// but panics on any errors.
func (l *Loader) MustGetTexture(name string, rend *sdl.Renderer) (*sdl.Surface, *sdl.Texture) {
	if s, t, err := l.GetTexture(name, rend); err != nil {
		panic(err)
	} else {
		return s, t
	}
}

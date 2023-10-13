package source

import (
	"github.com/pkg/errors"
)

type Options struct {
	Path         string
	Key          string
	AuthUser     string
	AuthPassword string
}

type NewSourceFunc func(opts Options) (Source, error)

var (
	ErrUnsupportedSource = errors.New("unsupported source")
	ErrNewFuncIsNil      = errors.New("source new func is nil")
)
var (
	supportedSources = map[string]NewSourceFunc{}
)

type Source interface {
	Read() ([]byte, error)
	String() string
}

func Register(name string, newFunc NewSourceFunc) {
	supportedSources[name] = newFunc
}

func New(typ string, opts Options) (Source, error) {
	newFunc, ok := supportedSources[typ]
	if !ok {
		return nil, ErrUnsupportedSource
	}
	if newFunc == nil {
		return nil, ErrNewFuncIsNil
	}
	return newFunc(opts)
}

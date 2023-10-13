package encoding

import "github.com/pkg/errors"

type NewEncodingFunc func() Encoder

var (
	supportedEncoding = map[string]NewEncodingFunc{}
)

var (
	ErrUnsupportedEncoding = errors.New("unsupported encoding")
	ErrNewFuncIsNil        = errors.New("encoding new func is nil")
)

type Encoder interface {
	Encode(interface{}) ([]byte, error)
	Decode([]byte, interface{}) error
	String() string
}

func Register(typ string, newFunc NewEncodingFunc) {
	supportedEncoding[typ] = newFunc
}

func New(typ string) (Encoder, error) {
	newFunc, ok := supportedEncoding[typ]
	if !ok {
		return nil, ErrUnsupportedEncoding
	}

	if newFunc == nil {
		return nil, ErrNewFuncIsNil
	}

	return newFunc(), nil
}

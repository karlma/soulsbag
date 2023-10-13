package encoding

import "github.com/pkg/errors"

type NewEncodingFunc func() Encoder

var (
	supportedEncoding = map[string]NewEncodingFunc{}
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
		return nil, errors.New("unsupported encoding")
	}

	if newFunc == nil {
		return nil, errors.New("encoding new func is nil")
	}

	return newFunc(), nil
}

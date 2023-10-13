package soulsbag

import (
	"github.com/karlma/soulsbag/encoding"
	_ "github.com/karlma/soulsbag/encoding/toml"
	_ "github.com/karlma/soulsbag/encoding/yaml"
	"github.com/karlma/soulsbag/source"
	_ "github.com/karlma/soulsbag/source/etcdv3"
	_ "github.com/karlma/soulsbag/source/file"
)

type SoulsBag struct {
	Source  source.Source
	Encoder encoding.Encoder
	Opts    source.Options
	Data    []byte
}

func New() *SoulsBag {
	return &SoulsBag{}
}

var sb *SoulsBag

func init() {
	sb = New()
}

func (s *SoulsBag) Init(sourceTyp, encodingTyp string, opts source.Options) error {
	src, err := source.New(sourceTyp, opts)
	if err != nil {
		return err
	}
	s.Source = src

	enc, err := encoding.New(encodingTyp)
	if err != nil {
		return err
	}
	s.Encoder = enc

	return nil
}

func Init(source, encoding string, opts source.Options) error {
	return sb.Init(source, encoding, opts)
}

func (s *SoulsBag) Read() error {
	data, err := s.Source.Read()
	if err != nil {
		return err
	}
	s.Data = data
	return nil
}
func Read() error {
	return sb.Read()
}

func (s *SoulsBag) Unmarshal(v interface{}) error {
	return s.Encoder.Decode(s.Data, v)
}

func Unmarshal(v interface{}) error {
	return sb.Unmarshal(v)
}
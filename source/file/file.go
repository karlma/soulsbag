package file

import (
	"errors"
	"os"

	"github.com/karlma/soulsbag/source"
)

type File struct {
	Path string
}

func (f File) Read() ([]byte, error) {
	if f.Path == "" {
		return nil, errors.New("file name is empty")
	}

	fileInfo, err := os.Stat(f.Path)
	if err != nil {
		return nil, err
	}
	if fileInfo.IsDir() {
		return nil, errors.New("file is directory")
	}

	data, err := os.ReadFile(f.Path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f File) String() string {
	return "file"
}

func NewSource(opts source.Options) (source.Source, error) {
	return File{Path: opts.Path}, nil
}

func init() {
	source.Register("file", NewSource)
}

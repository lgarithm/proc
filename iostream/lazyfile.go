package iostream

import (
	"fmt"
	"io"
	"os"
	"path"
)

func Open2Lazy(n1, n2 string) *combinedOut {
	return &combinedOut{
		outF: NewLazyFile(n1),
		errF: NewLazyFile(n2),
	}
}

type layzeFile struct {
	name string
	f    io.WriteCloser
}

func NewLazyFile(filename string) io.WriteCloser {
	return &layzeFile{name: filename}
}

func (f *layzeFile) Write(bs []byte) (int, error) {
	if f.f == nil {
		if err := f.create(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to create log file %s: %v", f.name, err)
			os.Stderr.Write(bs)
			return 0, err
		}
	}
	return f.f.Write(bs)
}

func (f *layzeFile) Close() error {
	if f.f != nil {
		return f.f.Close()
	}
	return nil
}

func (f *layzeFile) create() error {
	err := os.MkdirAll(path.Dir(f.name), os.ModePerm)
	if err != nil {
		return err
	}
	f.f, err = os.Create(f.name)
	return err
}

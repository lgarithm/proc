package iostream

import (
	"io"
	"os"
)

type combinedOut struct {
	outF, errF io.WriteCloser
}

func Open2(n1, n2 string) (*combinedOut, error) {
	outF, err := os.Create(n1)
	if err != nil {
		return nil, err
	}
	errF, err := os.Create(n2)
	if err != nil {
		return nil, err
	}
	return &combinedOut{
		outF: outF,
		errF: errF,
	}, nil
}

func (f *combinedOut) StdWriters() *StdWriters {
	return &StdWriters{
		Stdout: f.outF,
		Stderr: f.errF,
	}
}

func (f *combinedOut) Close() error {
	e1 := f.outF.Close()
	e2 := f.errF.Close()
	if e1 != nil {
		return e1
	}
	return e2
}

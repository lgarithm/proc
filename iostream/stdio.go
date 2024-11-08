package iostream

import (
	"io"
	"os"
	"sync"
)

var Std = StdWriters{
	Stdout: os.Stdout,
	Stderr: os.Stderr,
}

type StdReaders struct {
	Stdout io.Reader
	Stderr io.Reader
}

type StdWriters struct {
	Stdout io.Writer
	Stderr io.Writer
}

func (r *StdReaders) Stream(ws ...*StdWriters) interface{ Wait() } {
	var outs, errs []io.Writer
	for _, w := range ws {
		outs = append(outs, w.Stdout)
		errs = append(errs, w.Stderr)
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		Tee(r.Stdout, outs...)
		wg.Done()
	}()
	go func() {
		Tee(r.Stderr, errs...)
		wg.Done()
	}()
	return &wg
}

type WriterFunc func([]byte) (int, error)

func (w WriterFunc) Write(bs []byte) (int, error) { return w(bs) }

func MakeStdWriters(o, e WriterFunc) *StdWriters {
	return &StdWriters{
		Stdout: o,
		Stderr: e,
	}
}

// SaveFirstdWriter remembers the content of the first Write call
type SaveFirstdWriter struct {
	First string
}

func (w *SaveFirstdWriter) Write(bs []byte) (int, error) {
	if len(w.First) == 0 {
		w.First = string(bs)
	}
	return len(bs), nil
}

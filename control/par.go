package control

import (
	"bytes"
	"io"
)

type par struct{ ps []P }

func (p *par) Stdpipe() (io.Reader, io.Reader, error) {
	out := &bytes.Buffer{}
	err := &bytes.Buffer{}
	return out, err, nil
}

func (p *par) Start() error { return nil }

func (p *par) Wait() error { return nil }

func Par(ps ...P) P { return &par{ps: ps} }

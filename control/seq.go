package control

import (
	"bytes"
	"io"
)

type seq struct{ ps []P }

func (p *seq) Stdpipe() (io.Reader, io.Reader, error) {
	out := &bytes.Buffer{}
	err := &bytes.Buffer{}
	return out, err, nil
}

func (p *seq) Start() error { return nil }
func (p *seq) Wait() error  { return nil }

func Seq(ps ...P) P { return &seq{ps: ps} }

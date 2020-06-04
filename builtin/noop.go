package builtin

import (
	"bytes"
	"io"

	"github.com/lgarithm/proc-experimental/control"
)

type noop struct{}

func (p *noop) Stdpipe() (io.Reader, io.Reader, error) {
	out := &bytes.Buffer{}
	err := &bytes.Buffer{}
	return out, err, nil
}

func (p *noop) Start() error { return nil }

func (p *noop) Wait() error { return nil }

func Noop() control.P { return &noop{} }

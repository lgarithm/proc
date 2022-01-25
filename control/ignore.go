package control

import (
	"io"

	"github.com/lgarithm/proc-experimental/execution"
)

type ignore struct {
	p execution.P
}

func (p *ignore) Stdpipe() (io.Reader, io.Reader, error) {
	return p.p.Stdpipe()
}

func (p *ignore) Start() error {
	return p.p.Start()
}

func (p *ignore) Wait() error {
	p.p.Wait()
	return nil
}

func Ignore(q execution.P) execution.P {
	p := &ignore{p: q}
	return p
}

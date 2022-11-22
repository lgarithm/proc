package control

import (
	"io"
)

type ignore struct {
	p P
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

func Ignore(q P) P {
	p := &ignore{p: q}
	return p
}
